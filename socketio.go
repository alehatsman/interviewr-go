package main

import (
	"log"
	"net/http"

)

func main() {



	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
