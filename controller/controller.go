package controller

import (
	// "MarkMyLink/config"
	"MarkMyLink/model"
	"io"
	"io/ioutil"
	"net/http"
	//"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
	"log"
	"time"
	"encoding/json"
)
// Make sure that frontend sends the POST body in the json format.
//{
//"email": "biljithjayan@gmail.com",
//"name": "youtube",
//"link": "www.youtube.com",
//"viewcount": "1",
//"timestamp": "2020-11-10 23:00:00 +0000 UTC m=+0.000000000"
//}
func AddBookmark(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in.
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	_, err = model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		log.Fatal(err)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	var bookmark model.Bookmark
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &bookmark); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
	}

	if !model.CreateBookmark(bookmark) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
	http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}
// Make sure that frontend sends the POST body in the json format.
//{
//"email": "biljithjayan@gmail.com",
//"name": "youtube",
//"link": "www.youtube.com",
//"viewcount": "1",
//"timestamp": "2020-11-10 23:00:00 +0000 UTC m=+0.000000000"
//}
func UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in.
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	_, err = model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		log.Fatal(err)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	var bookmark model.Bookmark
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &bookmark); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if !model.UpdateBookmark(bookmark) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
	http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}
// Make sure that frontend sends the POST body in the json format.
//{
//"email": "biljithjayan@gmail.com",
//"name": "youtube",
//"link": "www.youtube.com",
//"viewcount": "1",
//"timestamp": "2020-11-10 23:00:00 +0000 UTC m=+0.000000000"
//}
func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	// Check if user is logged in.
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	_, err = model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		log.Fatal(err)
		return
	}
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	var bookmark model.Bookmark
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &bookmark); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	if !model.DeleteBookmark(bookmark) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
	http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}

type Token struct {
	Token string
}

func GetBookmarks(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "GET" {
	// 	http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 	return
	// }

	var jsonToken Token
	// Check if user is logged in.
	err := json.NewDecoder(r.Body).Decode(&jsonToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	sessionToken := jsonToken.Token
	session, err := model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	bm, err := model.FindBookmarks(session.Email)
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	bookmarksJson, err := json.Marshal(bm)
	if err != nil {
		// Deal with it
		log.Fatal(err)
	}
	w.Write(bookmarksJson)
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
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }
	// user := new(model.User)
	// decoder := schema.NewDecoder()

	// err = decoder.Decode(user, r.PostForm)
	// if err != nil {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }

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
	// http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Login")
	if r.Method != "POST" {
		// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		// return
		http.ServeFile(w, r, "templates/login.gohtml")
		return
	}

	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }
	// loginUser := new(model.User)
	// decoder := schema.NewDecoder()

	// err = decoder.Decode(loginUser, r.PostForm)
	// if err != nil {
	// 	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	// 	return
	// }
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Printf(sessionToken)
	json.NewEncoder(w).Encode(struct{
		Token string `json:"token"`
	}{Token: sessionToken})
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "session_token",
	// 	Value:   sessionToken,
	// 	Expires: session.CreatedAt.Add(3600 * time.Second),
	// })

	// http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}
