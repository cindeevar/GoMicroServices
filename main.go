package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	//implementation through handler
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)

	http.ListenAndServe(":9090", nil)
}
