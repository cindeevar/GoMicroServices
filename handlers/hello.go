// implementing the service through http handler
package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// adding logger as dependent object
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHttp(w http.ResponseWriter, r *http.Request) {

	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops error occurred", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s \n", d)

	fmt.Fprintf(w, "Hello %s", d)
}
