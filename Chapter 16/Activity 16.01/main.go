// adding a page counter to an HTML page
package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	counter int
	content string
	heading string
}

func (h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.counter++
	msg := fmt.Sprintf("<h1>%s</h1><p>%s</p><p>Views: %d</p>", h.heading, h.content, h.counter)
	w.Header().Add("Link", `<favicon.ico>; rel="icon"; hreflang="nofollow"`)
	// fmt.Println(w.Header().Get("Link"))
	w.Write([]byte(msg))
}

func main() {
	handler1 := PageWithCounter{0, "content 1", "Hello World"}
	handler2 := PageWithCounter{0, "content 2", "Chapter1"}
	handler3 := PageWithCounter{0, "content 3", "Chapter2"}
	http.Handle("/", &handler1)
	http.Handle("/chapter1", &handler2)
	http.Handle("/chapter2", &handler3)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
