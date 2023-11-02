package main

import (
	"log"
	"net/http"
    "fmt"
)

func main() {

    fmt.Println("Starting HTP Server on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("."))))
}
