package httpAndHttpsServer

import (
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
)

var (
	addr   = "127.0.0.1:8080"
	config *tls.Config
)

type Conn struct {
	net.Conn
	b byte
	e error
	f bool
}

func (c *Conn) Read(b []byte) (int, error) {
	if c.f {
		c.f = false
		b[0] = c.b
		if len(b) > 1 && c.e == nil {
			n, e := c.Conn.Read(b[1:])
			if e != nil {
				c.Conn.Close()
			}
			return n + 1, e
		} else {
			return 1, c.e
		}
	}
	return c.Conn.Read(b)
}

type SplitListener struct {
	net.Listener
}

func (l *SplitListener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	b := make([]byte, 1)
	_, err = c.Read(b)
	if err != nil {
		c.Close()
		if err != io.EOF {
			return nil, err
		}
	}

	con := &Conn{
		Conn: c,
		b:    b[0],
		e:    err,
		f:    true,
	}

	if b[0] == 22 {
		//log.Println("HTTPS")
		return tls.Server(con, config), nil
	}

	//log.Println("HTTP")
	return con, nil
}

func Run(addr, certFile, keyFile string, handler http.Handler)(err error){
	config = &tls.Config{}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	log.Fatal(http.Serve(
		&SplitListener{Listener: ln},
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.TLS == nil {
				//u := url.URL{
				//	Scheme:   "https",
				//	Opaque:   r.URL.Opaque,
				//	User:     r.URL.User,
				//	Host:     addr,
				//	Path:     r.URL.Path,
				//	RawQuery: r.URL.RawQuery,
				//	Fragment: r.URL.Fragment,
				//}
				//http.Redirect(w, r, u.String(), http.StatusMovedPermanently)
				http.DefaultServeMux.ServeHTTP(w, r)
			} else {
				http.DefaultServeMux.ServeHTTP(w, r)
			}
		})))
	return
}
