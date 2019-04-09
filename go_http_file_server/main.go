package main

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func main() {

	port := flag.String("port", "9090", "file server port")
	flag.Parse()
	//
	//if len(*root) == 0 {
	//	log.Fatalln("file server root directory not set")
	//}
	//
	//if !strings.HasPrefix(*root, "/") {
	//	log.Fatalln("file server root directory not begin with '/'")
	//}
	//
	//if !strings.HasSuffix(*root, "/") {
	//	log.Fatalln("file server root directory not end with '/'")
	//}
	root := "/"
	p, h := NewFileHandle(root)
	r := httprouter.New()
	r.GET(p, LogHandle(h))

	addr := ":" + string(*port)

	print("start listen ", addr)
	log.Fatalln(http.ListenAndServe(addr, r))
}

func NewFileHandle(path string) (string, httprouter.Handle) {
	return fmt.Sprintf("%s*files", path), func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.StripPrefix(path, http.FileServer(http.Dir(path))).ServeHTTP(w, r)
	}
}

func LogHandle(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		now := time.Now()
		handle(w, r, p)
		log.Printf("%s %s %s done in %v", r.RemoteAddr, r.Method, r.URL.Path, time.Since(now))
	}
}
