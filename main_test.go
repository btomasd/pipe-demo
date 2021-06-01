// package main_test

// import (
// 	"log"
// 	"net/http"
// )

// type ServerTest struct{}

// func (s *ServerTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message": "hello world"}`))
// }

// func main() {
// 	s := &ServerTest{}
// 	http.Handle("/", s)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Hello World!" {
		t.Fatal("Test fail")
	}
}

// func TestAdd(t *testing.T) {
// 	got := addition(1, 2)
// 	want := 3

// 	if got != want {
// 		t.Errorf("Oh no! Seems you can't add, you got %x", got)
// 	}

// }
