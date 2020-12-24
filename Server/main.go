package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	Routes()
}

//Handler new
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Another, %q", html.EscapeString(r.URL.Path))
}

//HandlerTwo new
func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Outra Rota que nem sei, %q", html.EscapeString(r.URL.Path))
}
