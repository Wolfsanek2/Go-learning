// requesting data from a web server and processing the response
package main

import "net/http"

type NamesHandler struct{}

func (h NamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := `{"names":["Electric", "Electric", "Boogalo", "Boogalo", "Boogalo"]}`
	w.Header().Add("Content-Type", "json")
	w.Write([]byte(msg))
}

func main() {
	handler := NamesHandler{}
	http.ListenAndServe(":8080", handler)
}
