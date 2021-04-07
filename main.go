package main


import (
	"MarkMyLink/controller"
	"net/http"
)

func main() {
	// Documentation for serving static files - 
	// https://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	//http.HandleFunc("/", index)
	// handles the URL 'localhost:8080/bookmarks
	http.HandleFunc(
		"/bookmarks", controller.Index)
	http.HandleFunc(
		"/signup", controller.Signup)
	http.HandleFunc(
		"/login", controller.Login)
	// this is where the web application will listen
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}