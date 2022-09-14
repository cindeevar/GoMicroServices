package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHttp(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}

	log.Printf("Data %s \n", d)

	fmt.Fprintf(w, "Hello %s", d)
}
