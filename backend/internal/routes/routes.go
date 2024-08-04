// internal/routes/routes.go
package routes

import (
	"manga-reader/backend/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

// InitRoutes инициализирует маршруты для API
func InitRoutes(router *gin.Engine) {
	// Swagger UI маршрут
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))

	// Основной маршрут
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Manga Reader API!")
	})

	// Маршруты для обработчиков
	router.GET("/search", handlers.SearchMangaHandler)
	router.GET("/scrap", handlers.ScrapMangaHandler)
	router.GET("/scrap/chapters", handlers.ScrapChaptersHandler)
	router.GET("/extract/images", handlers.ExtractImagesHandler)
}
