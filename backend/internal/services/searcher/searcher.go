package searcher

import (
	"fmt"
	models "manga-reader/backend/internal/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// SearchManga выполняет поиск по заданному запросу на указанном базовом URL и возвращает результат поиска.
//
// Параметры:
// - query: Строка запроса поиска.
// - baseURL: Базовый URL для выполнения поиска.
//
// Возвращает:
// - *SearchResult: Результат поиска, содержащий найденные манги.
// - error: Ошибка, если поиск не удался.
func SearchManga(query string) (*models.SearchResult, error) {
	baseURL := models.BaseURL
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	if strings.Contains(query, " ") {
		query = strings.ReplaceAll(query, " ", "+")
	}

	resp, err := http.Get(baseURL + "search?q=" + query)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при загрузке страницы поиска: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при парсинге HTML страницы поиска: %w", err)
	}

	var foundMangas []models.FoundManga

	doc.Find(".flex-container.row.align-items-start.justify-content-center .flex-item.card.mx-1.mx-md-2.mb-3.shadow-sm.rounded").Each(func(i int, s *goquery.Selection) {
		// Извлекаем URL манги
		href, exists := s.Find("a[href]").Attr("href")
		if exists {
			// Извлекаем URL обложки
			imageElem := s.Find("img.img-fluid.card-img-top")
			imageURL, exists := imageElem.Attr("src")
			if !exists || strings.HasPrefix(imageURL, "data:image") {
				imageURL, _ = imageElem.Attr("data-src")
			}
			// Извлекаем название манги
			title, _ := imageElem.Attr("title")

			// Создаем структуру FoundManga
			foundManga := models.FoundManga{
				URL:      models.BaseURL + href,
				Title:    title,
				ImageURL: imageURL,
			}
			// Добавляем в массив найденных манг
			foundMangas = append(foundMangas, foundManga)
		}
	})

	if len(foundMangas) == 0 {
		return nil, fmt.Errorf("Манга не найдена")
	}

	return &models.SearchResult{FoundMangas: foundMangas}, nil
}
