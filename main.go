package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func todos(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kiwkiw"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", users)
	mux.HandleFunc("/todos", todos)
	mux.HandleFunc("/", home)

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
