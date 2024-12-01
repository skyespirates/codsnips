package main

import (
	"codsnips.skyespirates.net/internal/models"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// example of running the server with a custom port
	// $ go run ./cmd/web -addr=":9999"		$$$$$$$$$$$$$$$$$$$$$
	// 2022/01/29 15:50:20 Starting server on :9999
	// Define the addr flag with a default value of ":4000" and short description
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "root:secret@tcp(127.0.0.1:3306)/codsnips?charset=utf8mb4&parseTime=True&loc=Local", "Mysql data source name")

	// Parse the command-line flags
	flag.Parse()

	// Leveled log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//cfg, e := mysql.ParseDSN(*dsn)
	//if e != nil {
	//	errorLog.Fatal("Failed to parse DSN: %v", e)
	//}
	//
	//conn, er := mysql.NewConnector(cfg)
	//if er != nil {
	//	errorLog.Fatal("Failed to create connector: %v", er)
	//}

	db, er := openDB(*dsn)
	if er != nil {
		errorLog.Fatal(er)
	}
	defer db.Close()

	//if err := db.Ping(); err != nil {
	//	errorLog.Println("Ping failed", err)
	//	os.Exit(1)
	//}

	infoLog.Println("Database connection is established")

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db}}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// log.Printf("Server is running on port %s", *addr)
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	// log.Fatal(err)
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
