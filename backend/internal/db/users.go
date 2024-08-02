package db

import (
	"database/sql"
	"manga-reader/backend/internal/models"
)

func CreateUser(db *sql.DB, user models.User) (int, error) {
	query := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`
	result, err := db.Exec(query, user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetUser(db *sql.DB, id int) (models.User, error) {
	query := `SELECT id, username, email, password_hash FROM users WHERE id = ?`
	var user models.User
	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	return user, err
}
