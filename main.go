package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
<<<<<<< HEAD

	_ "github.com/go-sql-driver/mysql"
=======
	"text/template"
>>>>>>> e62f5adcad85e6d23508b0ec9e8ac04d36624705
)

// Person is the structure for a person object
type Person struct {
	ID    string `json:"person_id"`
	FName string `json:"first_name"`
	LName string `json:"last_name"`
}

// Response is a list of person objects
type Response struct {
	People []Person `json:"result"`
}

var tmpl = template.Must(template.ParseGlob("form/*"))

// Index request data and displays the Index page
func Index(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8000/getpeople")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	person := Person{}
	res := []Person{}
	for i := 0; i < len(responseObject.People); i++ {
		person.ID = responseObject.People[i].ID
		person.FName = responseObject.People[i].FName
		person.LName = responseObject.People[i].LName
		res = append(res, person)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	fileServer := http.FileServer(http.Dir("./img/"))

	mux.Handle("/img/", http.StripPrefix("/img", fileServer))

<<<<<<< HEAD
	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
=======
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
	log.Println("Server started on port 8080")
>>>>>>> e62f5adcad85e6d23508b0ec9e8ac04d36624705
}
