package main

import (
	"AT_Elea/elea"
	"AT_Elea/src/server"
	"AT_Elea/util"
	"log"
	"runtime"
)

func init() {
	log.Println("The elea is in service")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// httpService()
	fileService()
}

func httpsService() {}

func httpService() {
	ip := util.LocalIP()
	Server := elea.Server{
		Addr:        ip + ":8080",
		Handle:      server.APIHandle(),
		Interceptor: &server.APIInterceptor{},
	}
	Server.ListenAndServer()
}

func fileService() {
	ip := util.LocalIP()
	Server := elea.Server{
		Addr:        ip + ":8090",
		Handle:      server.FileHandle(),
		Interceptor: &server.FileInterceptor{},
	}
	Server.ListenAndServer()
}

func socketService() {
	ip := util.LocalIP()
	elea.Aleph(ip, "9090")
}
