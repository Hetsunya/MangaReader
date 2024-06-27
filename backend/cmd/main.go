package main

import (
	"fmt"
	"manga-reader/backend/internal/services/imageextractor"
)

//imageextractor

const (
	baseURL = "https://mangapoisk.live" // Базовый URL сайта с мангой
)

func main() {
	// searchResult, _ := searcher.Search("повышение уровня", baseURL)

	// mangaURL := "https://mangapoisk.live/manga/i-have-90-billion-licking-gold?tab=chapters"

	// manga, err := scraper.Scrap(mangaURL)
	// if err != nil {
	// 	panic(err)
	// }

	// scraper.ScrapChapters(mangaURL, manga, baseURL)

	// scraper.PrintManga(manga)

	pages := imageextractor.ExtractImages("https://mangapoisk.live/manga/i-have-90-billion-licking-gold/chapter/1-1")
	for _, page := range pages {
		fmt.Println(page)
	}

	// TODO: Инициализация кэша

	// TODO: Запуск API

	// TODO: Настройка логирования

}
