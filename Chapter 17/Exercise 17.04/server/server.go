// uploading a file to a web server via a POST request
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadedFile, uploadedFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadedFile.Close()
	fileContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0600)
	w.Write([]byte(fmt.Sprintf("%s uploaded!", uploadedFileHeader.Filename)))
}

func main() {
	http.ListenAndServe(":8080", server{})
}
