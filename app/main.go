package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hi"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
