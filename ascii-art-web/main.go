package main

import (
	"01/ascii-art/common/functions"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	fmt.Println("http://localhost" + PORT)
	http.HandleFunc("/", functions.MainHandler)
	http.HandleFunc("/ascii-art-web", functions.HandleAscii)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
