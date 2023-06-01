package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, I love %s!", runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}