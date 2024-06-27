package imageextractor

import (
	"log"
	"net/http"
	"strings"

	"manga-reader/backend/internal/models"

	"github.com/PuerkitoBio/goquery"
)

// ExtractImages извлекает информацию о страницах манги с указанного URL.
// Функция возвращает слайс models.MangaPage, содержащий информацию о каждой странице манги.
func ExtractImages(url string) []models.MangaPage {
	var pages []models.MangaPage

	// HTTP GET запрос к указанному URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка при получении страницы: %s", resp.Status)
	}

	// Создание нового документа goquery из ответа
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Парсинг изображений
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists && strings.Contains(src, "pages") {
			page := models.MangaPage{
				ImageURL: src,
				// Дополнительные поля могут быть заполнены здесь
			}
			pages = append(pages, page)
		}
	})

	return pages
}
