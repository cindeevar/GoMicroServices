package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cindeevar/gomicroservices/handlers"
)

func main() {
	//implementation through handler
	l := log.New(os.Stdout, "")
	hh := handlers.NewHello(l)
	//concrete implementation
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
