package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello, World!")
}

func main() {
	http.HandleFunc("/hello", HelloWorld)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
