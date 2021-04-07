package tests

import (
	"MarkMyLink/controller"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login",controller.Login).Methods("POST")
	return router
}
func TestLogin(t *testing.T) {
 	request, _ := http.NewRequest("POST", "/login", nil)
 	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
 	assert.Equal(t, 200,response.Code, response.Code)

}


