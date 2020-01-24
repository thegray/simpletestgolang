package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprint(rw, "PBK simple response from golang\n")
		// log.Println("PBK simple response from golang")
	})
	port := os.Getenv("PORT")
	if port != "" {
		log.Println("Server started at:", port)
		http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	} else {
		log.Println("Env port missing, server started at :5000")
		http.ListenAndServe(":5000", nil)
	}
}
