package main

import (
	"fmt"
	"log"
	"manga-reader/backend/internal/routes"
	_ "manga-reader/docs"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Инициализация маршрутов
	routes.InitRoutes(router)

	// Запуск сервера
	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
