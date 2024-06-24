package main

import (
	"manga-reader/backend/internal/services/scraper"
	"manga-reader/backend/internal/services/searcher"
)

const (
	baseURL = "https://mangapoisk.live/search" // замените на базовый URL сайта с мангой
)

func main() {

	// TODO: Выполнение поиска

	mangaURL, _ := searcher.Search("повышение уровня", baseURL)

	// TODO: Доделать скрепер

	// mangaURL := "https://mangapoisk.live/manga/max-level-player"

	manga, _ := scraper.Scrap(mangaURL)

	scraper.PrintManga(manga)

	// TODO: Инициализация кэша

	// TODO: Запуск API

	// TODO: Настройка логирования

}
