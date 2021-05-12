package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("%s Error reading body: %v\n", r.RequestURI, err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		log.Printf("%s : %s\n", r.RequestURI, string(body))
		fmt.Fprintf(w, "ok")
	})

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
