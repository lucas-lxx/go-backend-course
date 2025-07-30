package main

import (
	"log"
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc(UrlBuilder(http.MethodGet, "/users"), api.getUserHandler)
	mux.HandleFunc(UrlBuilder(http.MethodPost, "/users"), api.postUserHandler)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
