// requesting data from a web server and processing the response
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Data struct {
	Names []string `json:names`
}

func getDataAndReturnResponse() Data {
	r, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	names := Data{}
	err = json.Unmarshal(data, &names)
	if err != nil {
		log.Fatal(err)
	}
	return names
}

func main() {
	data := getDataAndReturnResponse()
	names := data.Names
	counter := make(map[string]int)
	for _, name := range names {
		if _, ok := counter[name]; !ok {
			counter[name] = 1
		} else {
			counter[name]++
		}
	}
	for k, v := range counter {
		fmt.Printf("%s: %d\n", k, v)
	}
}
