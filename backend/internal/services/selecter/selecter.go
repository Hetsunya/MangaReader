package selecter

import (
	"fmt"
	"manga-reader/backend/internal/models"

	"golang.org/x/exp/slog"
)

// SelectManga позволяет пользователю выбрать мангу из списка найденных
func SelectManga(mangaSearchResult *models.SearchResult, log *slog.Logger) (string, error) {
	if mangaSearchResult == nil || len(mangaSearchResult.FoundMangas) == 0 {
		return "", fmt.Errorf("манга не найдена")
	}

	log.Info("Найдены следующие манги:")
	for i, manga := range mangaSearchResult.FoundMangas {
		log.Info("%d: %s\n", i+1, manga)
	}

	fmt.Print("Введите номер манги, которую хотите выбрать: ")
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 1 || choice > len(mangaSearchResult.FoundMangas) {
		return "", fmt.Errorf("неверный выбор: %v", err)
	}

	selectedMangaURL := mangaSearchResult.FoundMangas[choice-1].URL
	return selectedMangaURL, nil
}
