package handlers

import (
	"fmt"
	"net/http"
)

// Logger logs any http requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method + " " + r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
