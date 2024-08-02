package main

import (
	"fmt"
	"manga-reader/backend/internal/db"
	"manga-reader/backend/internal/lib/jsonutil"
	"manga-reader/backend/internal/models"
	"manga-reader/backend/internal/services/imageextractor"
	"manga-reader/backend/internal/services/scraper"
	"manga-reader/backend/internal/services/searcher"
	"os"

	"golang.org/x/exp/slog"
)

func setupLogger() *slog.Logger {
	var logger *slog.Logger

	logger = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	return logger
}

func main() {
	log := setupLogger()
	dbFilePath := "manga_reader.db"
	database := db.InitDB(dbFilePath)
	defer database.Close()

	user := models.User{
		Username:     "testuser",
		Email:        "testuser@example.com",
		PasswordHash: "hashedpassword",
	}

	userID, err := db.CreateUser(database, user)
	if err != nil {
		log.Error("Error creating user: %v", err)
		os.Exit(1)
	}

	log.Info("Created user with ID:", userID)

	mangaSearchResult, err := searcher.SearchManga("повышение уровня")
	if err != nil {
		log.Error("Ошибка при поиске манги: %v", err)
	}

	if mangaSearchResult == nil || len(mangaSearchResult.FoundMangas) == 0 {
		fmt.Println("Манга не найдена")
		return
	}

	var selectedMangaURL string

	log.Info("Найдены следующие манги:")
	for i, manga := range mangaSearchResult.FoundMangas {
		log.Info("%d: %s\n", i+1, manga)
	}

	fmt.Print("Введите номер манги, которую хотите выбрать: ")
	var choice int
	_, err = fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(mangaSearchResult.FoundMangas) {
		log.Error("Неверный выбор: %v", err)
	}

	selectedMangaURL = mangaSearchResult.FoundMangas[choice-1].URL

	log.Info(selectedMangaURL)

	manga, err := scraper.Scrap(selectedMangaURL)
	if err != nil {
		panic(err)
	}

	err = scraper.ScrapChapters(selectedMangaURL+models.ChapterParse, manga)
	if err != nil {
		panic(err)
	}

	scraper.PrintManga(manga)

	jsonString, _ := jsonutil.ToJSON(manga)
	fmt.Println(jsonString)

	pages := imageextractor.ExtractImages("https://mangapoisk.live/manga/i-have-90-billion-licking-gold/chapter/1-1")
	for _, page := range pages {
		fmt.Println(page)
	}

	mangaList := models.MangaList{
		UserID: userID,
		URL:    selectedMangaURL,
		Status: "читаю",
	}

	listID, err := db.CreateMangaList(database, mangaList)
	if err != nil {
		log.Error("Error creating manga list:", err)
		os.Exit(1)
	}

	log.Info("Created manga list with ID:", listID)

	lists, err := db.GetMangaListsByUserID(database, userID)
	if err != nil {
		log.Error("Error fetching manga lists:", err)
		os.Exit(1)
	}

	log.Info("Manga lists for user ID", userID, ":", lists)
}
