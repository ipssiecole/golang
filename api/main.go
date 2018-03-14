package main

import "api/server"

func main() {
	server.ListenAndServeTLS(":3000")
}
