package controller

import (
	"MarkMyLink/config"
	"MarkMyLink/model"

	"net/http"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
)

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
		log.Fatal(err)
	}
	user.Password = string(hashedPassword)
	if !model.CreateUser(user) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
}
