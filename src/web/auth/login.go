package auth

import (
	"AT_Elea/elea"
	"net/http"
)

func Login() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, elea.GenerateCookie())

		w.Write([]byte("Login API"))
	})
}
