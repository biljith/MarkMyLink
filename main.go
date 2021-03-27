package main


import (
	"MarkMyLink/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	// handles the URL 'localhost:8080/bookmarks
	http.HandleFunc(
		"/bookmarks", controller.Index)
	// this is where the web application will listen
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}