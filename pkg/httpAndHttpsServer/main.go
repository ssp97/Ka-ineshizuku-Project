package httpAndHttpsServer
//
//import (
//	"crypto/tls"
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"net/http"
//)
//
//var (
//	config *tls.Config
//)
//
//type Conn struct {
//	net.Conn
//	b byte
//	e error
//	f bool
//}
//
//func (c *Conn) Read(b []byte) (int, error) {
//	if len(b) == 0{
//		return c.Conn.Read(b)
//	}
//	if c.f {
//		c.f = false
//		b[0] = c.b
//		if len(b) > 1 && c.e == nil {
//			n, e := c.Conn.Read(b[1:])
//			if e != nil {
//				c.Conn.Close()
//			}
//			return n + 1, e
//		} else {
//			return 1, c.e
//		}
//	}
//	return c.Conn.Read(b)
//}
//
//type SplitListener struct {
//	net.Listener
//	config *tls.Config
//}
//
//func (l *SplitListener) Accept() (net.Conn, error) {
//	c, err := l.Listener.Accept()
//	if err != nil {
//		return nil, err
//	}
//
//	b := make([]byte, 1)
//	_, err = c.Read(b)
//	if err != nil {
//		c.Close()
//		if err != io.EOF {
//			return nil, err
//		}
//	}
//
//	con := &Conn{
//		Conn: c,
//		b:    b[0],
//		e:    err,
//		f:    true,
//	}
//
//	if b[0] == 22 {
//		fmt.Println("HTTPS")
//		return tls.Server(con, l.config), nil
//	}
//	fmt.Println("HTTP")
//	return con, nil
//}
//
//func Run(addr, certFile, keyFile string, handler http.Handler)(err error){
//	listener := SplitListener{}
//	listener.config = &tls.Config{}
//	if listener.config.NextProtos == nil {
//		listener.config.NextProtos = []string{"http/1.1"}
//	}
//	listener.config.Certificates = make([]tls.Certificate, 1)
//	listener.config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
//	if err != nil {
//		return
//	}
//
//	ln, err := net.Listen("tcp", addr)
//	listener.Listener = ln
//	if err != nil {
//		return
//	}
//	log.Fatal(http.Serve(
//		&listener,
//		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			if handler == nil{
//				http.DefaultServeMux.ServeHTTP(w, r)
//			}else{
//				handler.ServeHTTP(w,r)
//			}
//		})))
//	httputil.NewC
//	return
//}
