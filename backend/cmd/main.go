package main

import (
	"fmt"
	"log"
	"manga-reader/backend/internal/routes"
	_ "manga-reader/docs"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Создайте папку для логов, если она не существует
	logDir := "backend/logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Создайте файл для логов
	logFile, err := os.Create(fmt.Sprintf("%s/gin.log", logDir))
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()

	// Настройте логирование в файл
	gin.DefaultWriter = logFile
	gin.DefaultErrorWriter = logFile

	router := gin.Default()

	// Инициализация маршрутов
	routes.InitRoutes(router)

	// Запуск сервера
	fmt.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
