package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)

	return db
}

func createTables(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "username" TEXT NOT NULL,
        "email" TEXT NOT NULL,
        "password_hash" TEXT NOT NULL
    );`

	createMangaListTableSQL := `CREATE TABLE IF NOT EXISTS manga_lists (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "user_id" INTEGER NOT NULL,
        "url" TEXT NOT NULL,
        "status" TEXT
    );`

	createMangaTagTableSQL := `CREATE TABLE IF NOT EXISTS manga_tags (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "list_id" INTEGER NOT NULL,
        "tag" TEXT
    );`

	_, err := db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	_, err = db.Exec(createMangaListTableSQL)
	if err != nil {
		log.Fatalf("Error creating manga_lists table: %v", err)
	}

	_, err = db.Exec(createMangaTagTableSQL)
	if err != nil {
		log.Fatalf("Error creating manga_tags table: %v", err)
	}
}
