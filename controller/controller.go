package controller

import (
	"MarkMyLink/config"
	"MarkMyLink/model"
	"net/http"
	"fmt"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	//"log"
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	// logic to add bookmark
	//1. check for a logged in user and extract the user object
	//2. create in db
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bm, err := model.AllBookMarks()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("in controller")
	fmt.Println(bm)
	config.TPL.ExecuteTemplate(w, "index.gohtml", bm)
}


// Make sure that frontend sends the POST body in a correct format.
/*
 * {
 *   "email": "biljithjayan@gmail.com",
 *   "name": "Biljith",
 *   "password": "justabetterpassword"
 * }
 */
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)
	if !model.CreateUser(user) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	loginUser := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(loginUser); err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	storedUser, err := model.FindUser(loginUser.Email)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password),
												  []byte(loginUser.Password)); err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
	}
}
