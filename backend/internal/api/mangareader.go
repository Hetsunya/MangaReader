package mangareader

import (
	"fmt"
	"manga-reader/backend/internal/lib/jsonutil"
	"manga-reader/backend/internal/models"
	"manga-reader/backend/internal/services/scraper"
	"manga-reader/backend/internal/services/searcher"
)

// Exported function to be used in Android
func GetMangaJSON(query string) (string, error) {
	mangaSearchResult, err := searcher.SearchManga(query)
	if err != nil {
		return "", fmt.Errorf("Ошибка при поиске манги: %v", err)
	}

	if mangaSearchResult == nil || len(mangaSearchResult.FoundMangas) == 0 {
		return "", fmt.Errorf("Манга не найдена")
	}

	selectedMangaURL := mangaSearchResult.FoundMangas[0].URL

	manga, err := scraper.Scrap(selectedMangaURL)
	if err != nil {
		return "", fmt.Errorf("Ошибка при парсинге манги: %v", err)
	}

	err = scraper.ScrapChapters(selectedMangaURL+models.ChapterParse, manga)
	if err != nil {
		return "", fmt.Errorf("Ошибка при парсинге глав: %v", err)
	}

	jsonString, err := jsonutil.ToJSON(manga)
	if err != nil {
		return "", fmt.Errorf("Ошибка при конвертации в JSON: %v", err)
	}

	return jsonString, nil
}
