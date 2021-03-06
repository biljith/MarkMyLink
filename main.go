package main

import (
	"MarkMyLink/controller"
	"net/http"
	"os"
)


//helper function to check the return value for each amqp call




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
	 http.HandleFunc(
	 	"/updateBookmark", controller.UpdateBookmark)
	 http.HandleFunc(
	 	"/deleteBookmark", controller.DeleteBookmark)
	 http.HandleFunc(
	 	"/updateBookmarkVisit", controller.UpdateBookmarkVisit)
	// http.HandleFunc(
	// 	"/bookmarks", controller.Index)
	// this is where the web application will listen
	http.ListenAndServe(":" + port, nil)
}