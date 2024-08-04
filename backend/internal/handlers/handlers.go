package handlers

import (
	"fmt"
	"manga-reader/backend/internal/models"
	"manga-reader/backend/internal/services/imageextractor"
	"manga-reader/backend/internal/services/scraper"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// ExtractImagesHandler извлекает изображения манги по URL
// @Summary Извлечение изображений манги по URL
// @Description Извлекает изображения манги с указанного URL.
// @Tags Manga
// @Accept json
// @Produce json
// @Param url query string true "URL страницы манги"
// @Success 200 {array} models.MangaPage
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /extract/images [get]
func ExtractImagesHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL параметр 'url' обязателен"})
		return
	}

	pages, err := imageextractor.ExtractImages(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при извлечении изображений: %v", err)})
		return
	}

	c.JSON(http.StatusOK, pages)
}

// ScrapMangaHandler парсит информацию о манге по URL
// @Summary Парсинг информации о манге по URL
// @Description Парсит информацию о манге с указанного URL.
// @Tags Manga
// @Accept json
// @Produce json
// @Param url query string true "URL страницы манги"
// @Success 200 {object} models.Manga
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scrap [get]
func ScrapMangaHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL параметр 'url' обязателен"})
		return
	}

	manga, err := scraper.Scrap(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при скрапинге манги: %v", err)})
		return
	}

	c.JSON(http.StatusOK, manga)
}

// ScrapChaptersHandler парсит главы манги по URL
// @Summary Парсинг глав манги по URL
// @Description Парсит главы манги с указанного URL.
// @Tags Manga
// @Accept json
// @Produce json
// @Param url query string true "URL страницы манги"
// @Success 200 {array} models.Chapter
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scrap/chapters [get]
func ScrapChaptersHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL параметр 'url' обязателен"})
		return
	}

	var manga models.Manga
	err := scraper.ScrapChapters(url, &manga)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при скрапинге глав: %v", err)})
		return
	}

	c.JSON(http.StatusOK, manga.Chapters)
}

// SearchMangaHandler выполняет поиск манги по заданному запросу
// @Summary Поиск манги по запросу
// @Description Выполняет поиск манги по заданному запросу.
// @Tags Manga
// @Accept json
// @Produce json
// @Param q query string true "Поисковый запрос"
// @Success 200 {object} models.SearchResult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /search [get]
func SearchMangaHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query параметр 'q' обязателен"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при загрузке страницы поиска: %v", err)})
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Ошибка при парсинге HTML страницы поиска: %v", err)})
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
		c.JSON(http.StatusNotFound, gin.H{"message": "Манга не найдена"})
		return
	}

	result := models.SearchResult{FoundMangas: foundMangas}
	c.JSON(http.StatusOK, result)
}
