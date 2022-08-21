package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
