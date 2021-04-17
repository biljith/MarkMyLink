package main


import (
	"MarkMyLink/controller"
	"net/http"
	"os"
)

func main() {
	// Documentation for serving static files - 
	// https://www.alexedwards.net/blog/serving-static-sites-with-go
	http.HandleFunc(
		"/signup", controller.Signup)
	http.HandleFunc(
		"/login", controller.Login)
	http.HandleFunc(
		"/bookmarks", controller.GetBookmarks)
	fs := http.FileServer(http.Dir("./client/build"))
	http.Handle("/", fs)

	//http.HandleFunc("/", index)
	// handles the URL 'localhost:8080/bookmarks
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Commenting out below code. Uncomment after adding react UI.
	// http.HandleFunc(
	// 	"/bookmarks", controller.Index)
	// http.HandleFunc(
	// 	"/signup", controller.Signup)
	// http.HandleFunc(
	// 	"/login", controller.Login)
	http.HandleFunc(
		"/addBookmark", controller.AddBookmark)
	// http.HandleFunc(
	// 	"/updateBookmark", controller.UpdateBookmark)
	// http.HandleFunc(
	// 	"/deleteBookmark", controller.DeleteBookmark)
	// http.HandleFunc(
	// 	"/bookmarks", controller.Index)
	// this is where the web application will listen
	http.ListenAndServe(":" + port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/bookmarks", http.StatusSeeOther)
}