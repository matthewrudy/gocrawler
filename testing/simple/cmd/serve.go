package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/matthewrudy/gocrawler/testing/simple"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")

	handler := simple.Handler()
	log.Printf("Serving simple server on HTTP port: %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, handler))
}
