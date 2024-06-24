package searcher

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Search - функция для поиска манги на сайте
func Search(query string, baseURL string) (string, error) {
	// TODO: Проверка корректности baseURL
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}
	//TODO: Сделать + вместо проблема в query
	resp, err := http.Get(baseURL + "search?q=" + query)

	if err != nil {
		return "", fmt.Errorf("Ошибка при загрузке страницы поиска: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Ошибка при парсинге HTML страницы поиска: %w", err)
	}

	// TODO:  Используйте  правильный  селектор  для  сайта
	//  Например,  для  Mangalib  можно  использовать  ".search-results  .manga-item"
	var mangaURL string
	doc.Find(".search-results .manga-item a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		fmt.Println(href)
		if exists {
			mangaURL = href
			return
		}
	})

	if mangaURL == "" {
		return "", fmt.Errorf("Манга не найдена")
	}

	// TODO: Проверка правильности формата URL (например, начинается ли с baseURL)

	return mangaURL, nil
}
