package main

// sqlite3 installed on OS

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "sqlite/books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}
	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table books!")
	}
	statement.Exec()
	// Create
	statement, _ = db.Prepare("INSERT INTO books (name, author,isbn) VALUES (?,?,?)")
	statement.Exec("A table of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into database!")
	rows, _ := db.Query("SELECT id,name,author FROM books")
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID: %d, Book:%s, Author %s\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// Update
	statement, _ = db.Prepare("DELETE FROM books WHERE id=?")
	statement.Exec(1)
	log.Println("Succesfully deleted the book in the database!")

}
