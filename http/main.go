package main

import (
	"fmt"
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server\n"))
}

func (a *api) getIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET the index handler\n"))
}

func (a *api) postIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST the index handler\n"))
}

func (a *api) postAboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST the about handler\n"))
}

func UrlBuilder(method string, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}

func main() {
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()

	mux.HandleFunc(UrlBuilder(http.MethodGet, "/"), api.getIndexHandler)
	mux.HandleFunc(UrlBuilder(http.MethodPost, "/"), api.postIndexHandler)
	mux.HandleFunc(UrlBuilder(http.MethodPost, "/about"), api.postAboutHandler)
	if err := http.ListenAndServe(api.addr, mux); err != nil {
		log.Fatal(err)
	}
}
