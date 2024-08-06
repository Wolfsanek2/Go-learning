// personalized welcome
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	name, ok := v1["name"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
