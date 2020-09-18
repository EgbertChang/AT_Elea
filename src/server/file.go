package server

import (
	"AT_Elea/elea"
	"net/http"
)

// main函数中需要配置这个拦截器
type FileInterceptor struct{}

func (form *FileInterceptor) Intercept(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func FileHandle() elea.HandleSet {
	h := &elea.Handle{}
	FileHandlerRegister(h)
	return h
}

func FileHandlerRegister(h *elea.Handle) {
	h.Register("/file/", fileHandler)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {

}
