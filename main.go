package main

import (
  "123_hao_dai/elea"
  "123_hao_dai/src/server"
)

func main() {
  Server := elea.Server{
    Addr: ":8080",
    Handle: server.Handle(),
    Interceptor: &server.HttpInterceptor{},
  }
  Server.ListenAndServer()
}
