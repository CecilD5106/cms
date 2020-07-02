package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("Input URL")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	fileServer := http.FileServer(http.Dir("./img/"))

	mux.Handle("/img/", http.StripPrefix("/img", fileServer))

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
	log.Println("Server started on port 8080")
}
