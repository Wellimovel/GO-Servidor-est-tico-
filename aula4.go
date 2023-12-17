package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./estatico"))
	http.Handle("/", fs)
	log.Print("listening on: 3000...", nil)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
