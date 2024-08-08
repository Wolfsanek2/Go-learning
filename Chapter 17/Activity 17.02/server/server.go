// sending data to a web server and checking whether the data was received using POST and GET
package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type requestData struct {
	Name string `json:"name"`
}

type responseData struct {
	Ok bool `json:ok`
}

type getResponseData struct {
	Names []string `json:"names"`
}

type server struct {
	Data getResponseData
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGET(srv, w, r)
	case "POST":
		handlePOST(srv, w, r)
	}
}

func handlePOST(srv *server, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	jsonBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := requestData{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Fatal(err)
	}
	rspData := responseData{}
	if data.Name != "" {
		srv.Data.Names = append(srv.Data.Names, data.Name)
		rspData.Ok = true
	}
	jsonBytes, err = json.Marshal(&rspData)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonBytes)
}

func handleGET(srv *server, w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal(&srv.Data)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonBytes)
}

func main() {
	http.Handle("/", &server{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
