package main

import (
	"database/sql"
	"log"
	"net/http"

	pdb "sortedstartup.com/zero-to-release/db"
	"sortedstartup.com/zero-to-release/handlers"

	gHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	// Connect to SQLite database (file-based)
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	handlers.RegisterTaskHandlers(r, db)

	corsHandler := gHandlers.CORS(
		gHandlers.AllowedOrigins([]string{"*"}),
		gHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(r)

	// Run migrations
	err = pdb.MigrateDB("sqlite3", "sqlite3://./db.sqlite")
	if err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}
	log.Println("Migrating database done")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
