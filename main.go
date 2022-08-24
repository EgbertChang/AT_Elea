package main

import (
	"AT_Elea/elea"
	"AT_Elea/util"
	"log"
	"net/http"
	"runtime"
)

func init() {
	log.Println("🛰️The elea is in service")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	apiService()
	fileService()
	startServer()
}

func apiService() {
	http.HandleFunc("/path1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I am elea"))
	})
}

func fileService() {
	elea.FileService()
}

func socketService() {
	ip := util.LocalIP()
	elea.Aleph(ip, "9090")
}

func interceptorFunc(w http.ResponseWriter, r *http.Request) elea.SchedulerCode {
	return 0
}

func startServer() {
	ip := util.LocalIP()
	Server := elea.Server{
		Addr:            ip + ":8080",
		InterceptorFunc: interceptorFunc,
	}
	Server.ListenAndServe()
}
