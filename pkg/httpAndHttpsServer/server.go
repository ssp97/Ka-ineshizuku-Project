package httpAndHttpsServer
//
//import (
//	"context"
//	"net"
//	"net/http"
//	"sync"
//	"time"
//)
//
//type Server struct{
//	http.Server
//}
//
//type onceCloseListener struct {
//	net.Listener
//	once     sync.Once
//	closeErr error
//}
//
//var testHookServerServe func(*Server, net.Listener) // used if non-nil
//
//func Serve(l net.Listener, handler http.Handler) error {
//	srv := &Server{}
//	srv.Handler = handler
//	return srv.Serve(l)
//}
//
//func (srv *Server) Serve(l net.Listener) error {
//	if fn := testHookServerServe; fn != nil {
//		fn(srv, l) // call hook with unwrapped listener
//	}
//
//	origListener := l
//	l = &onceCloseListener{Listener: l}
//	defer l.Close()
//
//	if err := srv.setupHTTP2_Serve(); err != nil {
//		return err
//	}
//
//	if !srv.trackListener(&l, true) {
//		return http.ErrServerClosed
//	}
//	defer srv.trackListener(&l, false)
//
//	baseCtx := context.Background()
//	if srv.BaseContext != nil {
//		baseCtx = srv.BaseContext(origListener)
//		if baseCtx == nil {
//			panic("BaseContext returned a nil context")
//		}
//	}
//
//	var tempDelay time.Duration // how long to sleep on accept failure
//
//	ctx := context.WithValue(baseCtx, http.ServerContextKey, srv)
//	for {
//		rw, err := l.Accept()
//		if err != nil {
//			select {
//			case <-srv.getDoneChan():
//				return http.ErrServerClosed
//			default:
//			}
//			if ne, ok := err.(net.Error); ok && ne.Temporary() {
//				if tempDelay == 0 {
//					tempDelay = 5 * time.Millisecond
//				} else {
//					tempDelay *= 2
//				}
//				if max := 1 * time.Second; tempDelay > max {
//					tempDelay = max
//				}
//				srv.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
//				time.Sleep(tempDelay)
//				continue
//			}
//			return err
//		}
//		connCtx := ctx
//		if cc := srv.ConnContext; cc != nil {
//			connCtx = cc(connCtx, rw)
//			if connCtx == nil {
//				panic("ConnContext returned nil")
//			}
//		}
//		tempDelay = 0
//		c := srv.newConn(rw)
//		c.setState(c.rwc, http.StateNew, http.runHooks) // before Serve can return
//		go c.serve(connCtx)
//	}
//}
