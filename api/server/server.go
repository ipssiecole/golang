package server

import (
	"log"
	"net/http"
)

func ListenAndServeTLS(addr string) {
	key := "./server/ssl/server.key"
	crt := "./server/ssl/server.cert"
	if err := http.ListenAndServeTLS(addr, crt, key, nil); err != nil {
		log.Fatal(err)
	}
}