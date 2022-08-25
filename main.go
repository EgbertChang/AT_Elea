package main

import (
	"AT_Elea/elea"
	"AT_Elea/src/web/auth"
	"AT_Elea/util"
	"log"
	"net/http"
	"runtime"
)

func init() {
	log.Println("🛰️The name is Elea!")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	apiService()
	startServer()
}

func apiService() {
	auth.Login()
}

func fileService() {
	elea.FileService()
}

func socketService() {
	ip := util.LocalIP()
	elea.Aleph(ip, "10001")
}

func interceptorFunc(w http.ResponseWriter, r *http.Request) elea.SchedulerCode {
	return elea.Distribute
}

func startServer() {
	ip := util.LocalIP()
	Server := elea.Server{
		Addr:            ip + ":10000",
		InterceptorFunc: interceptorFunc,
	}
	Server.ListenAndServe()
}
