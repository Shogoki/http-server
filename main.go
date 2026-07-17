package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	addr := fmt.Sprintf("localhost:%d", *port)
	fmt.Printf("Starting HTTP Server on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("."))))
}
