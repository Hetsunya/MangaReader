package handlers

import (
	"manga-reader/backend/internal/models"
	"manga-reader/backend/internal/services/imageextractor"
	"manga-reader/backend/internal/services/scraper"
	"manga-reader/backend/internal/services/searcher"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Extract images from manga URL
// @Description Находит изображения 1 главы манги.
// @Tags manga
// @Accept json
// @Produce json
// @Param url query string true "Ссылка на главу"
// @Success 200 {array} models.MangaPage
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /extract/images [get]
func ExtractImagesHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "URL parameter 'url' is required"})
		return
	}

	pages, err := imageextractor.ExtractImages(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to extract images"})
		return
	}

	c.JSON(http.StatusOK, pages)
}

// @Summary Scrape manga information from URL
// @Description Берет инфу о манге с html страницы
// @Tags manga
// @Accept json
// @Produce json
// @Param url query string true "Ссылка на мангу"
// @Success 200 {object} models.Manga
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /scrap [get]
func ScrapMangaHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "URL parameter 'url' is required"})
		return
	}

	manga, err := scraper.Scrap(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to scrape manga information"})
		return
	}

	c.JSON(http.StatusOK, manga)
}

// @Summary Scrape manga chapters from URL
// @Description Берет инфу о главах манги с html страницы
// @Tags manga
// @Accept json
// @Produce json
// @Param url query string true "Ссылка на мангу + ?tab=chapters"
// @Success 200 {array} models.Chapter
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /scrap/chapters [get]
func ScrapChaptersHandler(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "URL parameter 'url' is required"})
		return
	}

	var manga models.Manga
	err := scraper.ScrapChapters(url, &manga)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to scrape chapters"})
		return
	}

	c.JSON(http.StatusOK, manga.Chapters)
}

// @Summary Search for manga by query
// @Description Находит мангу по запросу.
// @Tags manga
// @Accept json
// @Produce json
// @Param q query string true "Слова для поиска через +"
// @Success 200 {object} models.SearchResult
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /search [get]
func SearchMangaHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Query parameter 'q' is required"})
		return
	}

	result, err := searcher.SearchManga(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
