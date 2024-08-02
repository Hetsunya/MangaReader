package db

import (
	"database/sql"
	"manga-reader/backend/internal/models"
)

func CreateMangaList(db *sql.DB, list models.MangaList) (int, error) {
	query := `INSERT INTO manga_lists (name, url, status) VALUES (?, ?, ?)`
	result, err := db.Exec(query, list.Name, list.URL, list.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetMangaLists(db *sql.DB) ([]models.MangaList, error) {
	query := `SELECT id, name, url, status FROM manga_lists`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.MangaList
	for rows.Next() {
		var list models.MangaList
		err = rows.Scan(&list.ID, &list.Name, &list.URL, &list.Status)
		if err != nil {
			return nil, err
		}
		lists = append(lists, list)
	}

	return lists, nil
}
