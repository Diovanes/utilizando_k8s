package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":8000", nil)
}

func greeting(value string) string {
	return "<b>" + value + "</b>"
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", greeting(r.URL.Path[1:]))
}
