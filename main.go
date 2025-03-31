package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()

	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Database connected successfully!")

	// Create table
	if _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT
        )
    `); err != nil {
		log.Fatal(err)
	}

	// Insert a user
	if _, err = db.Exec("INSERT INTO users (name) VALUES (?)", "John Doe"); err != nil {
		log.Fatal(err)
	}

	// Query users
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print results
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User: %d, Name: %s\n", id, name)
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}).Methods(http.MethodPost)

	router.HandleFunc("/topics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}).Methods(http.MethodPost)

	http.ListenAndServe(":8080", router)
}
