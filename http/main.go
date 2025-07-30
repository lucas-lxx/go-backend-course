package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the server\n"))
}

func main() {
	s := &api{addr: ":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
