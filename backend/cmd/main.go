package main

import (
	"fmt"
	"log"
	"manga-reader/backend/internal/routes"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func initConfig() {
	// Установка путей для конфигурации
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Чтение конфигурационного файла
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func setupLogFile(logFilePath string) (*os.File, error) {
	// Создание директории для логов, если она не существует
	logDir := filepath.Dir(logFilePath)
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	// Открытие файла для логов
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open or create log file: %v", err)
	}

	// Настройка прав доступа к файлу логов
	if err := os.Chmod(logFilePath, 0644); err != nil {
		return nil, fmt.Errorf("failed to set log file permissions: %v", err)
	}

	return logFile, nil
}

func main() {
	// Инициализация конфигурации
	initConfig()

	// Получение пути к файлу логов из конфигурации
	logFilePath := viper.GetString("log.file")
	if logFilePath == "" {
		log.Fatalf("Log file path not specified in configuration")
	}

	// Настройка логирования в файл
	logFile, err := setupLogFile(logFilePath)
	if err != nil {
		log.Fatalf("Could not setup log file: %v", err)
	}
	defer logFile.Close()

	gin.DefaultWriter = logFile
	gin.DefaultErrorWriter = logFile

	// Создание маршрутизатора
	router := gin.Default()

	// Инициализация маршрутов
	routes.InitRoutes(router)

	// Запуск сервера
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080" // Значение по умолчанию
	}
	fmt.Printf("Starting server on :%s\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
