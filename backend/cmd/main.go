package main

import (
	"fmt"
	"log"
	_ "manga-reader/backend/cmd/docs"
	"manga-reader/backend/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Manga Reader API
// @version 1.0
// @description API для поиска манги и получения информации о манге
// @host localhost:8080
// @basePath /
func main() {
	router := gin.Default()

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

	// Запуск сервера
	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
