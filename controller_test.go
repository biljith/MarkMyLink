package main

import (
	"MarkMyLink/controller"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login",controller.Login).Methods("POST")
	router.HandleFunc("/signup",controller.Signup).Methods("POST")
	return router
}

func TestSignUpBadRequest(t *testing.T) {
	request, _ := http.NewRequest("POST", "/signup", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t, 400,response.Code, "Expected 400 Badrequest")
}

func TestSignUpStatusConflict(t *testing.T) {
	body := "{\"email\": \"biljithjayan@gmail.com\",\"name\": \"Biljith\",\"password\": \"justabetterpassword\" }"
	request, _ := http.NewRequest("POST", "/signup", strings.NewReader(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t, 409,response.Code, "Expected 409 Status Conflict")
}

func TestLoginBadRequest(t *testing.T) {
 	request, _ := http.NewRequest("POST", "/login", nil)
 	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
 	assert.Equal(t, 400,response.Code, "Expected 400 Badrequest")
}

func TestLoginSeeOther(t *testing.T) {
	body := "{\"email\": \"biljithjayan@gmail.com\",\"name\": \"Biljith\",\"password\": \"justabetterpassword\" }"
	request, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t, 303,response.Code, "Expected 303 See Other")
}

func TestLoginUserUnauthorized(t *testing.T) {
	body := "{\"email\": \"biljithjayan@gmail.com\",\"name\": \"Biljith\",\"password\": \"justabetterpassword\" }"
	request, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t, 401,response.Code, "Expected 401 Unauthorized")
}

func TestLoginInvalidSession(t *testing.T) {
	body := "{\"email\": \"biljithjayan@gmail.com\",\"name\": \"Biljith\",\"password\": \"justabetterpassword\" }"
	request, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	response := httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	request2, _ := http.NewRequest("POST", "/login", strings.NewReader(body))
	response2 := httptest.NewRecorder()
	Router().ServeHTTP(response2,request2)
	assert.Equal(t, 500,response.Code, "Expected 500 Unauthorized")
}

