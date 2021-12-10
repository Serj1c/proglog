package main

import (
	"log"

	"github.com/serj1c/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8060")
	log.Fatal(srv.ListenAndServe())
}
