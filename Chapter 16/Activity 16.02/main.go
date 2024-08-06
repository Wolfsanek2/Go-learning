// external template
package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Visitor struct {
	Name string
}

type Hello struct {
	tmpl *template.Template
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := Visitor{}
	query := r.URL.Query()
	name, ok := query["name"]
	if ok {
		v.Name = strings.Join(name, ",")
	}
	err := h.tmpl.Execute(w, v)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Panic(err)
	}
	handler := Hello{tmpl}
	http.Handle("/", &handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
