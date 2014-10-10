package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Get port
	port := flag.String("port", "1979", "Port number to listen on")

	// Get path
	path := flag.String("path", ".", "Path to publish.  Default is current directory")

	flag.Parse()

	url := "localhost:" + *port
	fmt.Printf("Publishing \"%v\" on \"%v\"", *path, url)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*path))))
	if err := http.ListenAndServe(":"+string(*port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
