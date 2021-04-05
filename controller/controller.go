package controller

import (
	"MarkMyLink/config"
	"MarkMyLink/model"
	"net/http"
	"fmt"
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
	fmt.Println("in controller")
	fmt.Println(bm)
	config.TPL.ExecuteTemplate(w, "index.gohtml", bm)
}
