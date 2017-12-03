package main

import (
	"log"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() {
	os.Remove("./foo.db")

	sqlStmt := ` 
		CREATE TABLE IF NOT EXISTS user (ID INTEGER NOT NULL PRIMARY KEY, name TEXT);
		CREATE TABLE IF NOT EXISTS device (ID INTEGER NOT NULL PRIMARY KEY, user_ID INTEGER, name TEXT, FOREIGN KEY(user_ID) REFERENCES user(ID));
		CREATE TABLE IF NOT EXISTS subscription (ID INTEGER NOT NULL PRIMARY KEY, user_ID INTEGER, name TEXT, FOREIGN KEY(user_ID) REFERENCES user(ID));

		INSERT INTO user VALUES(1, "Zim"); 
		INSERT INTO user VALUES(2, "Gir"); 

		INSERT INTO device VALUES(1, 1, "Laptop");
		INSERT INTO device VALUES(2, 2, "Phone");
		
		INSERT INTO subscription VALUES(1, 1, "http://podcast.com");
		INSERT INTO subscription VALUES(2, 2, "http://podcast.com");
		INSERT INTO subscription VALUES(NULL, 2, "http://zerocast.com");
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		defer db.Close()
		return
	}
}

func FindTodo(id int) string {
	stmt, err := db.Prepare("select todo from todo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var todo string
	err = stmt.QueryRow(id).Scan(&todo)
	// if err != nil {
	// 	// log.Fatal(err)
	// 	fmt.Println(err)
	// }
	return todo
}

