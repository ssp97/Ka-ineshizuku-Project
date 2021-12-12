package driver

import (
	"errors"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

var (
	nullResponse = zero.APIResponse{}
	json         = jsoniter.ConfigFastest
)

type WSServer struct {
	//Port        int // ws连接地址
	//AccessToken string
	Server      websocket.Upgrader
	handler     func([]byte, zero.APICaller)
}

type WSConn struct {
	seq         uint64
	selfID      int64
	Server      *WSServer
	conn        *websocket.Conn
	mu          sync.Mutex // 写锁
	seqMap      seqSyncMap
}

func NewWebSocketServer() *WSServer{
	return &WSServer{
		//Port:         port,
		//AccessToken: accessToken,
		Server : websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (ws *WSServer) Connect() {
	// nothing
}

func (ws *WSServer) SelfID() int64{
	return 0
}

func (ws *WSServer)websocketServer(w http.ResponseWriter, r *http.Request){
	wsConn := WSConn{
		Server: ws,
		selfID: 0,
	}

	c, err := ws.Server.Upgrade(w, r, nil)
	if err != nil{
		log.Error("ws connect:", err)
	}
	defer c.Close()
	wsConn.conn = c

	log.Infof("接受websocket连接:%v", c.RemoteAddr())

	go func() {
		rsp, err := wsConn.CallApi(zero.APIRequest{
			Action: "get_login_info",
			Params: nil,
		})

		if err != nil{
			log.Warn("获取基础信息失败")
			return
		}
		wsConn.selfID = rsp.Data.Get("user_id").Int()
		zero.APICallers.Store(wsConn.selfID, &wsConn) // 添加Caller到 APICaller list...
		log.Infof("基础信息获取成功:%v", c.RemoteAddr())
	}()

	for {
		t, payload, err := wsConn.conn.ReadMessage()
		if err != nil { // reconnect
			if wsConn.selfID > 0{
				zero.APICallers.Delete(wsConn.selfID)   // 退出后则要删除掉
			}
			log.Warnf("websocket连接断开:%v", c.RemoteAddr())
			return
		}

		if t == websocket.TextMessage {
			rsp := gjson.Parse(helper.BytesToString(payload))
			if rsp.Get("echo").Exists() { // 存在echo字段，是api调用的返回
				log.Debug("接收到API调用返回: ", strings.TrimSpace(helper.BytesToString(payload)))
				go func(rsp gjson.Result) {
					if c, ok := wsConn.seqMap.LoadAndDelete(rsp.Get("echo").Uint()); ok {
						c <- zero.APIResponse{ // 发送api调用响应
							Status:  rsp.Get("status").String(),
							Data:    rsp.Get("data"),
							Msg:     rsp.Get("msg").Str,
							Wording: rsp.Get("wording").Str,
							RetCode: rsp.Get("retcode").Int(),
							Echo:    rsp.Get("echo").Uint(),
						}
						close(c) // channel only use once
					}
				}(rsp)
			} else {
				if rsp.Get("meta_event_type").Str != "heartbeat" { // 忽略心跳事件
					log.Debug("接收到事件: ", helper.BytesToString(payload))
				}
				go wsConn.Server.handler(payload, &wsConn)
			}
		}
	}
}

func (ws *WSServer) Listen(handler func([]byte, zero.APICaller)){
	ws.handler = handler
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", ws.websocketServer)
	http.HandleFunc("/api/bot", ws.websocketServer)
	//log.Infof("start server on port %d", ws.Port)
	//err := http.ListenAndServe(fmt.Sprintf(":%d", ws.Port),mux)
	//if err!=nil{
	//	log.Error(err)
	//}
}

func (ws *WSConn) nextSeq() uint64 {
	return atomic.AddUint64(&ws.seq, 1)
}

// CallApi 发送ws请求
func (ws *WSConn) CallApi(req zero.APIRequest) (zero.APIResponse, error) {
	ch := make(chan zero.APIResponse)
	req.Echo = ws.nextSeq()
	ws.seqMap.Store(req.Echo, ch)
	data, err := json.Marshal(req)
	if err != nil {
		return nullResponse, err
	}

	// send message
	ws.mu.Lock() // websocket write is not goroutine safe
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	ws.mu.Unlock()
	if err != nil {
		log.Warn("向WebsocketServer发送API请求失败: ", err.Error())
		return nullResponse, err
	}

	log.Debug("向服务器发送请求: ", helper.BytesToString(data))
	select { // 等待数据返回
	case rsp, ok := <-ch:
		if !ok {
			return nullResponse, errors.New("channel closed")
		}
		return rsp, nil
	case <-time.After(30 * time.Second):
		ws.seqMap.Delete(req.Echo)
		return nullResponse, errors.New("timed out")
	}
}
