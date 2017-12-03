package main

import (
    "log"
    "net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	InitDB()
	
    router := NewRouter()

    log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
