package server

import (
	"net/http"
	"bytemurmur.com/server/handler"
	"log"
	"sync"
	"bytemurmur.com/server/components/router"
)

type HTTPServer struct{
	stopChan  	chan struct{}
	lock      	*sync.Mutex
	srv    		*http.Server
}

func NewHTTPServer() (*HTTPServer, error) {

	//router.Register("/test", HandlerRequest)

	srv := &HTTPServer{
		stopChan: 	make(chan struct{}, 1),
		lock: 		&sync.Mutex{},
		srv: 		&http.Server{Addr: ":4621"},
	}

	srv.srv.Handler = &router.HTTPHandler{}

	router.Register("/", &handler.DashboardHandler{})

	go srv.server()

	return srv, nil
}

func (server *HTTPServer) Stop() {
	server.lock.Lock()
	defer server.lock.Unlock()
	server.srv.Shutdown(nil)
}

func (server *HTTPServer) server () {
	http.HandleFunc("/", handler.DashboardIndex)
	err := server.srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
