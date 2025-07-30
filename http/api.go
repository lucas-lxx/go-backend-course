package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users = []User{}

type api struct {
	addr string
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server\n"))
}

func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *api) postUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload User
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u := User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	users = append(users, u)

	w.WriteHeader(http.StatusCreated)
}

func UrlBuilder(method string, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}
