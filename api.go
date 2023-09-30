package main

//Lessons
/*
1. don't camelCase when in differnt files it doesn't have "scope" or something to access
2. package name SHOULD be main (why need to check)
3. use go run . and not specific file else everything doesn't compile and throws error
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func MakeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAdder string
	store 		Storage
}

func NewAPIServer(listenAdder string, store Storage) *APIServer {
	return &APIServer{
		listenAdder: listenAdder,
		store:		 store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", MakeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", MakeHTTPHandleFunc(s.handleGetAccount))

	log.Println("JSON API server running on port : ", s.listenAdder)

	http.ListenAndServe(s.listenAdder, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}

	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}

	// if r.Method == "UPDATE"{
	// 	return s.handleGetAccount(w, r)
	// }

	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	return WriteJSON(w, http.StatusOK, vars)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
