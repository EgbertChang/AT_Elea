package elea

import (
	"log"
	"net/http"
	"time"
)

type SchedulerCode int

const (
	Distribute SchedulerCode = iota
	SessionExpiration
)

type interceptorFunc func(http.ResponseWriter, *http.Request) SchedulerCode

// 调度器类用于处理所有http请求的顶层逻辑
type scheduler struct {
	InterceptorFunc interceptorFunc
}

func (sc *scheduler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if sc.InterceptorFunc == nil || sc.InterceptorFunc(w, r) == Distribute {
		http.DefaultServeMux.ServeHTTP(w, r)
	}
}

type Server struct {
	Addr            string
	InterceptorFunc interceptorFunc
}

func (s *Server) ListenAndServe() {
	server := http.Server{
		Addr:           s.Addr,
		Handler:        &scheduler{InterceptorFunc: s.InterceptorFunc},
		ReadTimeout:    3 * time.Minute,
		WriteTimeout:   3 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Elea Server is listening on " + s.Addr)
	log.Fatalln(server.ListenAndServe())
}
