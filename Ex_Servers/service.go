package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá\n")
}

func cabecalhos(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalhos", cabecalhos)

	http.ListenAndServe(":8080", nil)
}
