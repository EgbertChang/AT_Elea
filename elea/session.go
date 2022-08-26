package elea

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const expirationDuration = 1 * time.Hour

func generateSessionId() string {
	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Int63()
	return strconv.FormatInt(randNum, 10)
}

func GenerateCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "ELEAUID",
		Value:   generateSessionId(),
		Expires: time.Now().Add(expirationDuration),
		Path:    "/",
		Domain:  LocalIP(),
	}
}
