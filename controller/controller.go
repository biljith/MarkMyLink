package controller

import (
	"MarkMyLink/config"
	"MarkMyLink/model"
	"net/http"
	"fmt"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"log"
	"time"
)

func AddBookmark(w http.ResponseWriter, r *http.Request) {
	// logic to add bookmark
	//1. check for a logged in user and extract the user object
	//2. create in db
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in.
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
			log.Fatal(err)
			return
		}
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	session, err := model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		log.Fatal(err)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bm, err := model.FindBookmarks(session.Email)
	log.Printf("%v", bm)
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
		// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		// return
		http.ServeFile(w, r, "templates/signup.gohtml")
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
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		// return
		http.ServeFile(w, r, "templates/login.gohtml")
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

	sessionToken := uuid.NewString()

	session := &model.Session{}
	session.Email = loginUser.Email
	session.Token = sessionToken
	session.CreatedAt = time.Now()
	if !model.CreateSession(session) {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: session.CreatedAt.Add(3600 * time.Second),
	})
}
