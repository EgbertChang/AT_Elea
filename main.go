package main

import (
  "123_hao_dai/util"
  "123_hao_dai/elea"
  "123_hao_dai/src/server"
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
