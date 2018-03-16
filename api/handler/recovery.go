package handler

import (
	"log"
	"net/http"
)

// Recovery catch any panics.
func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				SendJSONError(w, "", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
