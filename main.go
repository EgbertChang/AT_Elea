package main

import (
  "AT_Elea/elea"
  "AT_Elea/src/server"
)

func main() {
  Server := elea.Server{
    Addr: ":8080",
    Handle: server.Handle(),
    Interceptor: &server.HttpInterceptor{},
  }
  Server.ListenAndServer()
}
