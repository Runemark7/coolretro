package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID    int
	Title string
}

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	if _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT NOT NULL
        )
    `); err != nil {
		log.Fatal(err)
	}

	// Insert two users
	if _, err = db.Exec("INSERT INTO tasks (title) VALUES (?), (?)", "Learn Go", "Build HTMX app"); err != nil {
		log.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handleTasks(db, w, r)
	})

	log.Println("Server running on localhost:8080")

	http.ListenAndServe(":8080", nil)
}

func handleTasks(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" { // Check if the request comes from HTMX
		tasks := fetchTasks(db)

		tmpl := template.Must(template.New("tasks").Parse(`
			{{range .}}
			<tr>
				<td>{{.ID}}</td>
				<td>{{.Title}}</td>
			</tr>
			{{end}}
        `))
		tmpl.Execute(w, tasks)
		return
	}

	http.Error(w, "Not Found", http.StatusNotFound)
}

func fetchTasks(db *sql.DB) []Task {
	rows, err := db.Query("SELECT id, title FROM tasks")
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title); err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, t)
	}
	return tasks
}
