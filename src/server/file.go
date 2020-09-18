package server

import (
	"AT_Elea/elea"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	path := r.URL.Path
	steps := strings.Split(path, "/")
	fileName := steps[len(steps)-1]
	file, err := os.Open("src/assets/" + fileName)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`<body align=center><h2>404 Not Found</h2><hr>Elea 0.0.1</body>`))
		return
	}
	fileBytes, _ := ioutil.ReadAll(file)
	w.Write(fileBytes)
}
