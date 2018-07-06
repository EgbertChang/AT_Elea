package main

import (
  "AT_Elea/util"
  "AT_Elea/elea"
  "AT_Elea/src/server"
)

func main() {
  ip := util.LocalIP()
  Server := elea.Server{
    Addr: ip + ":8080",
    Handle: server.Handle(),
    Interceptor: &server.HttpInterceptor{},
  }
  Server.ListenAndServer()
}
