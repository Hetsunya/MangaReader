package handlers

import (
	"encoding/json"
	"fmt"
	"manga-reader/backend/internal/models"
	"manga-reader/backend/internal/services/imageextractor"
	"manga-reader/backend/internal/services/scraper"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractImagesHandler извлекает изображения манги по URL
func ExtractImagesHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL параметр 'url' обязателен", http.StatusBadRequest)
		return
	}

	pages, err := imageextractor.ExtractImages(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при извлечении изображений: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pages)
}

// ScrapMangaHandler парсит информацию о манге по URL
func ScrapMangaHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL параметр 'url' обязателен", http.StatusBadRequest)
		return
	}

	manga, err := scraper.Scrap(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при скрапинге манги: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manga)
}

// ScrapChaptersHandler парсит главы манги по URL
func ScrapChaptersHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL параметр 'url' обязателен", http.StatusBadRequest)
		return
	}

	var manga models.Manga
	err := scraper.ScrapChapters(url, &manga)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при скрапинге глав: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manga.Chapters)
}

// SearchMangaHandler выполняет поиск манги по заданному запросу.
func SearchMangaHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	baseURL := models.BaseURL
	if !strings.HasSuffix(baseURL, "/") {
		baseURL += "/"
	}

	if strings.Contains(query, " ") {
		query = strings.ReplaceAll(query, " ", "+")
	}

	resp, err := http.Get(baseURL + "search?q=" + query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при загрузке страницы поиска: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при парсинге HTML страницы поиска: %v", err), http.StatusInternalServerError)
		return
	}

	var foundMangas []models.FoundManga

	doc.Find(".flex-container.row.align-items-start.justify-content-center .flex-item.card.mx-1.mx-md-2.mb-3.shadow-sm.rounded").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Find("a[href]").Attr("href")
		if exists {
			title := s.Find("h2.entry-title").Text()
			foundManga := models.FoundManga{
				URL:   models.BaseURL + href,
				Title: title,
			}
			foundMangas = append(foundMangas, foundManga)
		}
	})

	if len(foundMangas) == 0 {
		http.Error(w, "Манга не найдена", http.StatusNotFound)
		return
	}

	result := models.SearchResult{FoundMangas: foundMangas}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
