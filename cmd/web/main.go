package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// example of running the server with a custom port
	// $ go run ./cmd/web -addr=":9999"		$$$$$$$$$$$$$$$$$$$$$
	// 2022/01/29 15:50:20 Starting server on :9999
	// Define the addr flag with a default value of ":4000" and short description
	addr := flag.String("addr", ":8080", "HTTP network address")
	// Parse the command-line flags
	flag.Parse()

	// Leveled log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// log.Printf("Server is running on port %s", *addr)
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	// log.Fatal(err)
	errorLog.Fatal(err)
}
