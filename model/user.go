package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// User defines the user data in the database
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

// GetUsers gets a list of all users
func GetUsers() []User {
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

	return users
}

func GetUser(ID int) (*User, error) {
	for _, c := range categories {
		if c.ID == ID {
			return &c, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Category with ID %v not found", ID))
}
