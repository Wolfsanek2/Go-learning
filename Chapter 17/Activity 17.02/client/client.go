// sending data to a web server and checking whether the data was received using POST and GET
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func addNameAndParseResponse(name string) (bool, error) {
	client := http.DefaultClient
	sendData := requestData{Name: name}
	jsonBytes, err := json.Marshal(&sendData)
	if err != nil {
		return false, err
	}
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return false, err
	}
	response, err := client.Do(request)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	jsonBytes, err = io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	responseData := responseData{}
	err = json.Unmarshal(jsonBytes, &responseData)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getDataAndParseResponse() ([]string, error) {
	client := http.DefaultClient
	request, err := http.NewRequest("GET", "http://127.0.0.1:8080", nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	jsonBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := getResponseData{}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Names, nil
}

func main() {
	ok, err := addNameAndParseResponse("Electric")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
	ok, err = addNameAndParseResponse("Boogalo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)

	names, err := getDataAndParseResponse()
	fmt.Println(names)
}
