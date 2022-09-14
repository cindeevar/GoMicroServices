package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}

		log.Printf("Data %s \n", d)

		fmt.Fprintf(w, "Hello %s", d)
	})
	http.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Println("GoodBye World")
	})
	http.ListenAndServe(":9000", nil)
}
