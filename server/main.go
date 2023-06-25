package main

import (
	"log"
	"net/http"
	"webart"
)

/*
This is the main function that starts the server. It is in this directory that typing (go run .)
will the server actually start listening and serving
*/

func main() {
	http.HandleFunc("/", webart.Handler)  
	http.HandleFunc("/w", webart.Index)              //* listens for the "/" append to the localhost URL and and handles is using the webart.Handler function
	http.HandleFunc("/ascii-art", webart.Gen_ASCII)     //* listens for the "/ascii-art" append to the localhost URL and and handles is using the webart.ascii-art function
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err) //* logging any errors in the terminal
	}
}