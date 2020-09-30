package main

import (
	"log"

	"github.com/bohdan-chechin/proglog/internals/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
