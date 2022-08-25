package auth

import "net/http"

func Login() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login API"))
	})
}
