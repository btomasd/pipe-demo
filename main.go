// package main

// import (
// 	"log"
// 	"net/http"
// )

// type Server struct{}

// func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(`{"message": "hello world"}`))
// }

// func main() {
// 	s := &Server{}
// 	http.Handle("/", s)
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

package main

import "fmt"

func helloworld() string {
	return "Hello World!!!!!!!"
}

// func addition(a, b int) (res int) {
// 	return a + b
// }

func main() {
	fmt.Println(helloworld())
}
