package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "cgdavis"
	dbPass := "DzftXvz$eR7VpY^h"
	dbServer := "tcp(172.18.105.227:3306)"
	dbName := "people"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbServer+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

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
	log.Println("Server started on port 8080")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
