package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req: %#v", r)
		next.ServeHTTP(w, r)
	})
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s", name)
}

func NumProvider(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	num := vars["num"]
	fmt.Fprintf(w, "Your number is: %s", num)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", HelloWorld)
	r.HandleFunc("/numis/{num}", NumProvider)
	r.Use(LogMiddleware)
	http.Handle("/", r)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
