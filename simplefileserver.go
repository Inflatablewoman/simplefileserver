package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/user"
)

func main() {

	// Get port
	port := flag.String("port", "1979", "Port number to listen on")

	// Work out the current users home directory
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Get path
	path := flag.String("path", usr.HomeDir, "Path to publish")

	flag.Parse()

	url := "localhost:" + *port
	fmt.Printf("Publishing \"%v\" on \"%v\"", *path, url)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*path))))
	if err := http.ListenAndServe(":"+string(*port), nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
