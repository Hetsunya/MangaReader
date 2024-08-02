package db

import (
	"database/sql"
	"log"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	createTableQueries := []string{
		`CREATE TABLE IF NOT EXISTS manga_lists (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			url TEXT UNIQUE,
			status TEXT CHECK(status IN ('читаю', 'в планах', 'брошено', 'готово', 'любимое'))
		);`,
		`CREATE TABLE IF NOT EXISTS manga_tags (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			list_id INTEGER,
			tag TEXT,
			FOREIGN KEY (list_id) REFERENCES manga_lists(id)
		);`,
	}

	for _, query := range createTableQueries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
