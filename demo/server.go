package main

import (
	"flag"
	"log"
	"net/http"
)

var output *string

func init() {
	output = flag.String("output", "output", "defines the output folder")
	flag.Parse()
}

func main() {
	fs := http.FileServer(http.Dir(*output))
	log.Fatal(http.ListenAndServe(":8080", fs))
}
