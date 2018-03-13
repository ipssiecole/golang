package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my homepage."))
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About us.")
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("My Middleware")
		next.ServeHTTP(w, r)
	})
}

func main() {
	handler := http.HandlerFunc(home)
	http.Handle("/", Middleware(handler))
	http.HandleFunc("/about", about)

	http.ListenAndServe(":3000", nil)
}
