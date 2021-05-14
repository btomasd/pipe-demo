package main_test

import (
	"log"
	"net/http"
)

type ServerTest struct{}

func (s *ServerTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &ServerTest{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
