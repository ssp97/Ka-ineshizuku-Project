package httpAndHttpsServer
//
//import (
//	"fmt"
//	"io"
//	"log"
//	"net"
//	"net/http"
//	"os"
//	"path"
//)
//
//const (
//	SOCK_HTTP = "sockHttp.sock"
//	SOCK_HTTPS = "sockHttps.sock"
//)
//
//func Run(addr, certFile, keyFile string, handler http.Handler)(err error) {
//	dir,_ := os.Getwd()
//	sockHttp := path.Join(dir, "data", SOCK_HTTP)
//	sockHttps := path.Join(dir, "data", SOCK_HTTPS)
//	os.Remove(sockHttp)
//	os.Remove(sockHttps)
//
//	go RunHttpSock(sockHttp, handler)
//	//go RunHttpsSock(sockHttps, certFile, keyFile, handler)
//
//	go RunHttpAndHttpsMux(addr, sockHttp, sockHttps)
//	return
//}
//
//func RunHttpSock(file string, handler http.Handler){
//	listener, err := net.Listen("unixpacket", file)
//	if err != nil{
//		log.Panic(err)
//		return
//	}
//	log.Fatal(http.Serve(
//		listener,
//		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			if handler == nil{
//				http.DefaultServeMux.ServeHTTP(w, r)
//			}else{
//				handler.ServeHTTP(w,r)
//			}
//		})))
//}
//
//func RunHttpsSock(file, certFile, keyFile string, handler http.Handler){
//	listener, err := net.Listen("unixpacket", file)
//	if err != nil{
//		log.Panic(err)
//		return
//	}
//	log.Fatal(http.ServeTLS(
//		listener,
//		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			if handler == nil{
//				http.DefaultServeMux.ServeHTTP(w, r)
//			}else{
//				handler.ServeHTTP(w,r)
//			}
//		}),
//		certFile,
//		keyFile,
//		))
//}
//
//func RunHttpAndHttpsMux(addr, httpSock, httpsSock string){
//	listener, err := net.Listen("tcp", addr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		go func(local net.Conn) {
//			sniff := make([]byte, 1)
//			n, err := local.Read(sniff)
//			if err != nil && err != io.EOF {
//				log.Fatal(err)
//			}
//			var file string
//			if sniff[0] == 0x22 {
//				file = httpsSock
//			} else {
//				file = httpSock
//			}
//
//			remote, err := net.Dial("unix", file)
//			if err != nil {
//				log.Fatal(err)
//			}
//			go io.Copy(local, remote)
//			go func() {
//				remote.Write(sniff[:n])
//				io.Copy(remote, local)
//			}()
//		}(conn)
//	}
//}
//
//func forwarder() {
//
//	listener, err := net.Listen("tcp", ":12345")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for {
//		local, err := listener.Accept()
//		if err != nil {
//			log.Fatal(err)
//		}
//		sniff := make([]byte, 3)
//		n, err := local.Read(sniff)
//		if err != nil && err != io.EOF {
//			log.Fatal(err)
//		}
//		s := string(sniff[:n])
//
//		var addr string
//		if s == "GET"[:n] || s == "POS"[:n] {
//			addr = ":12346"
//		} else {
//			addr = ":12347"
//		}
//
//		remote, err := net.Dial("tcp", addr)
//		if err != nil {
//			log.Fatal(err)
//		}
//		go io.Copy(local, remote)
//		go func() {
//			remote.Write(sniff[:n])
//			io.Copy(remote, local)
//		}()
//	}
//}
//
//func Http(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "HTTP")
//		handler.ServeHTTP(w, r)})
//}
//
//func Https(handler http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, "HTTPS")
//		handler.ServeHTTP(w, r)})
//}
//
//func hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Hello world!")
//}