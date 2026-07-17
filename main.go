package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.1.0"

func main() {
	host := flag.String("host", "localhost", "address to listen on (use 0.0.0.0 to listen on all interfaces)")
	port := flag.Int("port", 8080, "port to listen on")
	showVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println("http-server", version)
		os.Exit(0)
	}

	addr := fmt.Sprintf("%s:%d", *host, *port)
	fmt.Printf("Starting HTTP Server on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("."))))
}
