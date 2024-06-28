package main

import (
	"fmt"
	"manga-reader/backend/internal/lib/jsonutil"
	"manga-reader/backend/internal/models"
	"manga-reader/backend/internal/services/scraper"
)

func main() {
	// searchResult, _ := searcher.Search("повышение уровня", baseURL)
	//"https://mangapoisk.live/manga/i-have-90-billion-licking-gold?tab=chapters"

	mangaURL := "https://mangapoisk.live/manga/i-have-90-billion-licking-gold"

	manga, err := scraper.Scrap(mangaURL)
	if err != nil {
		panic(err)
	}

	scraper.ScrapChapters(mangaURL+models.ChapterParse, manga)

	scraper.PrintManga(manga)

	jsonString, _ := jsonutil.ToJSON(manga)

	fmt.Println(jsonString)
	// pages := imageextractor.ExtractImages("https://mangapoisk.live/manga/i-have-90-billion-licking-gold/chapter/1-1")
	// for _, page := range pages {
	// 	fmt.Println(page)
	// }

	// TODO: Инициализация кэша

	// TODO: Запуск API

	// TODO: Настройка логирования

}
