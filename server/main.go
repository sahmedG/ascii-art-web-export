package main

import (
	"fmt"
	"log"
	"net/http"
	"webart"
)

/*
This is the main function that starts the server. It is in this directory that typing (go run .)
will the server actually start listening and serving
*/

func main() {
	http.HandleFunc("/", webart.Handler)             //* listens for the "/" append to the localhost URL and and handles is using the webart.Handler function
	http.HandleFunc("/export", webart.ExportHandler) //* listens for the "/export" append to the localhost URL and and handles is using the webart.ExportHandler function
	//http.HandleFunc("/export.txt", webart.ExportHandler)
	http.HandleFunc("/ascii-art", webart.Gen_ASCII)  //* listens for the "/ascii-art" append to the localhost URL and and handles is using the webart.Gen_ASCII function
	fmt.Println("Server runing on port 8080, goto localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err) //* logging any errors in the terminal
	}
}
