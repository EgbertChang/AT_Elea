package server

import (
	"AT_Elea/elea"
	"net/http"
)

// main函数中需要配置这个拦截器
type APIInterceptor struct{}

func (form *APIInterceptor) Intercept(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func APIHandle() elea.HandleSet {
	h := &elea.Handle{}
	APIHandlerRegister(h)
	return h
}

func APIHandlerRegister(h *elea.Handle) {
	h.Register("/api/path1", APIPath1)
	h.Register("/api/path2", APIPath2)
	h.Register("/api/path3", APIPath3)
}

func APIPath1(w http.ResponseWriter, r *http.Request) {
	ua := r.Header.Get("user-agent")
	w.Write([]byte(ua))
}

func APIPath2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Path2"))
}

func APIPath3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<body><h1>Hello Path3</h1></body>"))
}
