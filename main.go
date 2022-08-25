package main

import (
	"fmt"
	"log"
	"net/http"
)

type LogMiddleware struct {
}

func (l *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("req: %#v", r)
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	lm := LogMiddleware{}
	http.Handle("/", &lm)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
