package elea

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func FileService() {
	http.HandleFunc("/file/", fileHandler)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	path := r.URL.Path
	steps := strings.Split(path, "/")
	fileName := steps[len(steps)-1]
	file, err := os.Open("src/assets/" + fileName)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`<body align=center><h2>404 Not Found</h2><hr>elea 0.0.1</body>`))
		return
	}
	fileBytes, _ := ioutil.ReadAll(file)
	w.Write(fileBytes)
}
