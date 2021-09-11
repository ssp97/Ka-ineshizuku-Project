package httpAndHttpsServer

import (
	"crypto/rand"
	"crypto/tls"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
	"strings"
)

func serveHTTP1(l net.Listener, handler http.Handler) {
	s := &http.Server{
		Handler: handler,
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		panic(err)
	}
}

func serveHTTPS(l net.Listener,handler http.Handler,certFile, keyFile string) {
	// Load certificates.
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		Rand:         rand.Reader,
	}

	// Create TLS listener.
	tlsl := tls.NewListener(l, config)

	// Serve HTTP over TLS.
	serveHTTP1(tlsl, handler)
}


func Run(addr, certFile, keyFile string, handler http.Handler)(err error){
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error(err)
	}

	m := cmux.New(l)
	httpl := m.Match(cmux.HTTP1Fast())
	tlsl := m.Match(cmux.Any())

	go serveHTTP1(httpl, handler)
	go serveHTTPS(tlsl,handler, certFile, keyFile)

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		panic(err)
	}

	return
}