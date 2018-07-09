package main

import (
	"AT_Elea/elea"
	"AT_Elea/src/server"
	"AT_Elea/util"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	ip := util.LocalIP()
	Server := elea.Server{
		Addr:        ip + ":8080",
		Handle:      server.Handle(),
		Interceptor: &server.HttpInterceptor{},
	}
	Server.ListenAndServer()
}
