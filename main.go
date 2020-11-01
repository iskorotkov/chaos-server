package main

import (
	"flag"
	"fmt"
	"net/http"
)

var counter = 0
var port = flag.Int("port", 8811, "Port to listen on")

func main() {
	flag.Parse()

	http.HandleFunc("/counter", handleCounter)
	fmt.Println("Server is starting")
	err := http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
	if err != nil {
		fmt.Printf("Can't listen to port %v\n", port)
		panic(err)
	}
}

func handleCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := fmt.Fprint(w, counter)
		if err != nil {
			panic(err)
		}
	} else if r.Method == "POST" {
		counter++
	} else {
		panic("Unsupported method")
	}
}
