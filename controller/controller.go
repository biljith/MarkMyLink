package controller

import (
	// "MarkMyLink/config"
	"MarkMyLink/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	//"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)
// Make sure that frontend sends the POST body in the json format.
//{
//"email": "biljithjayan@gmail.com",
//"name": "youtube",
//"link": "www.youtube.com",
//"viewcount": "1",
//"timestamp": "2020-11-10 23:00:00 +0000 UTC m=+0.000000000"
//}


type Token struct {
	Token string
}
type TokenAndLink struct {
	Token string `bson:"token"`
	Link string  `bson:"link"`
	Name string  `bson:"name"`
}
type LinkPreviewResponse struct {
	Title string `bson:"title"`
	Description string `bson:"description"`
	Image string `bson:"image"`
	Url string `bson:"url"`
}
func AddBookmark(w http.ResponseWriter, r *http.Request) {
	var tl TokenAndLink
	// Check if user is logged in.
	err := json.NewDecoder(r.Body).Decode(&tl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	sessionToken := tl.Token
	session, err := model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	var bookmark model.Bookmark
	bookmark.Name = tl.Name
	bookmark.Link = tl.Link
	bookmark.Email = session.Email
	bookmark.Viewcount = 1
	dt := time.Now()
	bookmark.Timestamp = dt.String()
	url := "http://api.linkpreview.net/?key=%s&q=%s"
	linkPreviewAPIKey := os.Getenv("LINK_PREVIEW_API_KEY")
	res, err := http.Get(fmt.Sprintf(url, linkPreviewAPIKey, bookmark.Link))
	if err != nil {
		log.Printf("couldn't find link preview")
	}
	defer res.Body.Close()
	linkPreview := &LinkPreviewResponse{}
	json.NewDecoder(res.Body).Decode(linkPreview)
	bookmark.Image = linkPreview.Image
	bookmark.Description = linkPreview.Description
	if !model.CreateBookmark(bookmark) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
}
func UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	var tl TokenAndLink
	// Check if user is logged in.
	err := json.NewDecoder(r.Body).Decode(&tl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessionToken := tl.Token
	session, err := model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	var bookmark model.Bookmark
	bookmark.Name = tl.Name
	bookmark.Link = tl.Link
	bookmark.Email = session.Email
	dt := time.Now()
	bookmark.Timestamp = dt.String()
	bookmark.Viewcount = 1
	url := "http://api.linkpreview.net/?key=%s&q=%s"
	linkPreviewAPIKey := os.Getenv("LINK_PREVIEW_API_KEY")
	res, err := http.Get(fmt.Sprintf(url, linkPreviewAPIKey, bookmark.Link))
	if err != nil {
		log.Printf("couldn't find link preview")
	}
	defer res.Body.Close()
	linkPreview := &LinkPreviewResponse{}
	json.NewDecoder(res.Body).Decode(linkPreview)
	bookmark.Image = linkPreview.Image
	bookmark.Description = linkPreview.Description
	if !model.UpdateBookmark(bookmark) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
}

func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	var tl TokenAndLink
	// Check if user is logged in.
	err := json.NewDecoder(r.Body).Decode(&tl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sessionToken := tl.Token
	session, err := model.FindSession(sessionToken)
	if err != nil {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}
	var bookmark model.Bookmark
	bookmark.Link = tl.Link
	bookmark.Email = session.Email

	if !model.DeleteBookmark(bookmark) {
		http.Error(w, http.StatusText(409), http.StatusConflict)
		return
	}
}

func GetBookmarks(w http.ResponseWriter, r *http.Request) {

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
