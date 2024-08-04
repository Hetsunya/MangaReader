// package main

// import (
// 	"fmt"
// 	"manga-reader/backend/internal/db"
// 	"manga-reader/backend/internal/lib/jsonutil"
// 	"manga-reader/backend/internal/models"
// 	"manga-reader/backend/internal/services/scraper"
// 	"manga-reader/backend/internal/services/searcher"
// 	"os"

// 	_ "github.com/mattn/go-sqlite3"

// 	"golang.org/x/exp/slog"
// )

// func setupLogger() *slog.Logger {
// 	var logger *slog.Logger

// 	logger = slog.New(
// 		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
// 	)

// 	return logger
// }

// func main() {
// 	log := setupLogger()
// 	dbFilePath := "manga_reader.db"
// 	database := db.InitDB(dbFilePath)
// 	defer database.Close()

// 	mangaSearchResult, err := searcher.SearchManga("повышение уровня")
// 	if err != nil {
// 		log.Error("Ошибка при поиске манги: %v", err)
// 		return
// 	}

// 	if mangaSearchResult == nil || len(mangaSearchResult.FoundMangas) == 0 {
// 		fmt.Println("Манга не найдена")
// 		return
// 	}

// 	log.Info("Найдены следующие манги:")
// 	for i, manga := range mangaSearchResult.FoundMangas {
// 		log.Info("%d: %s\n", i+1, manga)
// 	}

// 	fmt.Print("Введите номер манги, которую хотите выбрать: ")
// 	var choice int
// 	_, err = fmt.Scan(&choice)
// 	if err != nil || choice < 1 || choice > len(mangaSearchResult.FoundMangas) {
// 		log.Error("Неверный выбор: %v", err)
// 		return
// 	}

// 	selectedManga := mangaSearchResult.FoundMangas[choice-1]
// 	selectedMangaURL := selectedManga.URL

// 	log.Info(selectedMangaURL)

// 	manga, err := scraper.Scrap(selectedMangaURL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = scraper.ScrapChapters(selectedMangaURL+models.ChapterParse, manga)
// 	if err != nil {
// 		panic(err)
// 	}

// 	scraper.PrintManga(manga)

// 	jsonString, _ := jsonutil.ToJSON(manga)
// 	fmt.Println(jsonString)

// 	exists, err := db.CheckMangaExists(database, selectedMangaURL)
// 	if err != nil {
// 		log.Error("Error checking manga existence:", err)
// 		return
// 	}
// 	if exists {
// 		log.Info("Манга уже существует в базе данных")
// 		return
// 	}

// 	mangaList := models.MangaList{
// 		Name:   manga.Title[0],
// 		URL:    selectedMangaURL,
// 		Status: "читаю",
// 	}

// 	listID, err := db.CreateMangaList(database, mangaList)
// 	if err != nil {
// 		log.Error("Error creating manga list:", err)
// 		return
// 	}

// 	log.Info("Created manga list with ID:", listID)

// 	lists, err := db.GetMangaLists(database)
// 	if err != nil {
// 		log.Error("Error fetching manga lists:", err)
// 		return
// 	}

// 	log.Info("Manga lists:", lists)
// }

package main

import (
	"fmt"
	"log"
	"manga-reader/backend/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Manga Reader API!")
	})

	http.HandleFunc("/search", handlers.SearchMangaHandler)
	http.HandleFunc("/scrap", handlers.ScrapMangaHandler)             // Новый маршрут для скрапа манги
	http.HandleFunc("/scrap/chapters", handlers.ScrapChaptersHandler) // Новый маршрут для скрапа глав
	http.HandleFunc("/extract/images", handlers.ExtractImagesHandler) // Новый маршрут для извлечения изображений

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
