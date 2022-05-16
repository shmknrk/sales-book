package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
