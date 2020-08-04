package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// User is the structure for a user object
type User struct {
	ID              string `json:"user_id"`
	UserName        string `json:"user_name"`
	UserEmail       string `json:"user_email"`
	FName           string `json:"user_first_name"`
	LName           string `json:"user_last_name"`
	Password        string `json:"password"`
	PasswordChange  string `json:"password_change"`
	PasswordExpired string `json:"password_expired"`
	LastLogon       string `json:"last_logon"`
	AccountLocked   string `json:"account_locked"`
}

// Response is a list of person objects
type Response struct {
	Users []User `json:"result"`
}

var tmpl = template.Must(template.ParseGlob("form/*"))

const pwdSalt = "5X!cw*V9byQ3@v9ct!Cv&Lq4X#m8Ci27pteC&n7$6Nq4VUgkqzP5woC7oK!5stXH*zJ9W86E@GpgCjP78jGoWsrA@jkMWPkF&avNKi6grWs@$bMr7pg&3hyf"

// UserList request data and displays the user list page
func UserList(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://uapi:8000/v1/getusers")
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

	user := User{}
	users := []User{}
	for i := 0; i < len(responseObject.Users); i++ {
		user.ID = responseObject.Users[i].ID
		user.UserName = responseObject.Users[i].UserName
		user.UserEmail = responseObject.Users[i].UserEmail
		user.FName = responseObject.Users[i].FName
		user.LName = responseObject.Users[i].LName
		user.Password = responseObject.Users[i].Password
		user.PasswordChange = responseObject.Users[i].PasswordChange
		user.PasswordExpired = responseObject.Users[i].PasswordExpired
		user.LastLogon = responseObject.Users[i].LastLogon
		user.AccountLocked = responseObject.Users[i].AccountLocked
		users = append(users, user)
	}
	tmpl.ExecuteTemplate(w, "UserList", users)
}

//Index displays the Index page
func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

//NewUser opens the new user page in edit mode
func NewUser(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewUser", nil)
}

//Logon opens the logon page in edit mode
func Logon(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Logon", nil)
}

//VerifyLogon checks the password submitted by the user
func VerifyLogon(w http.ResponseWriter, r *http.Request) {
	//Get User from api

	//Create password hash from data on the form

	//Compare password hash

	//If fail loop through the 3 times

	//If successful move to Index page
}

//InsertUser saves the data from the NewUser page to the database
func InsertUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		useremail := r.FormValue("useremail")
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		pwd := r.FormValue("password")
		pwdExp := r.FormValue("passwordexpired")
		pwdChg := r.FormValue("passwordchange")
		lstLgn := r.FormValue("lastlogon")
		acctLkd := r.FormValue("accountlocked")

		// Set password expired and account locked information to
		// false if checkbox is not checked
		if acctLkd == "" {
			acctLkd = "0"
		}
		if pwdExp == "" {
			pwdExp = "0"
		}

		// Creates a hash of the password to store in the database
		hasher := sha512.New()
		hasher.Write([]byte(username))
		hasher.Write([]byte(pwdSalt))
		hasher.Write([]byte(pwd))
		pwd = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

		jsonData := map[string]string{"user_name": username, "user_email": useremail, "user_first_name": fname, "user_last_name": lname, "password": pwd, "password_change": pwdChg, "password_expired": pwdExp, "last_logon": lstLgn, "account_locked": acctLkd}
		jsonValue, _ := json.Marshal(jsonData)
		res, err := http.Post("http://uapi:8000/v1/createuser", "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			panic(err.Error())
		} else {
			data, _ := ioutil.ReadAll(res.Body)
			fmt.Println(string(data))
		}
	}
	http.Redirect(w, r, "/UserList", 301)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/NewUser", NewUser)
	mux.HandleFunc("/insertuser", InsertUser)
	mux.HandleFunc("/UserList", UserList)

	fileServer := http.FileServer(http.Dir("./img/"))

	mux.Handle("/img/", http.StripPrefix("/img", fileServer))

	log.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
