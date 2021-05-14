package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hokaccha/go-prettyjson"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("%s Error reading body: %v\n", r.RequestURI, err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}
		if prettyHeader, err := prettyjson.Marshal(r.Header); err != nil {
			log.Printf("%s : %s\n\n", r.RequestURI, string(body))
		} else {
			log.Printf("%s : %s\n%s\n\n", r.RequestURI, string(body), string(prettyHeader))
		}
		fmt.Fprintf(w, "ok")
	})

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
