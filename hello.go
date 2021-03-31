package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Println("listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
