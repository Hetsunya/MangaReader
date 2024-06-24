package scraper

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Manga представляет информацию о манге
type Manga struct {
	Title       []string
	Description string
	Genre       []string
	Chapters    []Chapter
}

// Chapter представляет информацию о главе манги
type Chapter struct {
	Title string
	Date  string
	Link  string
}

// Scrap - функция для парсинга страницы манги
func Scrap(url string) (*Manga, error) {
	c := colly.NewCollector()

	// Установка лимитов для запросов
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,               // Запросов в секунду
		Delay:       1 * time.Second, // Можно уменьшить, если допустимо
	})

	manga := &Manga{}

	// Получение заголовков манги
	c.OnHTML(".card-header.p-0, .text-sm.overflow-hidden.inline.manga-alt-name", func(e *colly.HTMLElement) {
		titles := strings.Split(e.Text, "/")
		for _, title := range titles {
			manga.Title = append(manga.Title, strings.TrimSpace(title))
		}
	})

	// Получение жанров манги
	c.OnHTML(".badge.variant-soft-tertiary.mb-1.mr-1", func(e *colly.HTMLElement) {
		manga.Genre = append(manga.Genre, e.Text)
	})

	// Получение описания манги
	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		manga.Description = e.Attr("content")
	})

	// Получение информации о главах
	c.OnHTML("li.flex.justify-between.card.variant-soft-surface.mb-3.p-2.md\\:px-3", func(e *colly.HTMLElement) {
		chapter := Chapter{
			Title: strings.TrimSpace(e.ChildText("span.chapter-title")),
			Date:  strings.TrimSpace(e.ChildText("span.text-surface-400.text-right.whitespace-nowrap.text-sm")),
			Link:  strings.TrimSpace(e.ChildAttr("a", "href")),
		}
		manga.Chapters = append(manga.Chapters, chapter)
	})

	// TODO: Сделать парс глав
	// c.OnHTML(".li.justify-between.card.variant-soft-surface.mb-3.p-2.md:px-3", func(e *colly.HTMLElement) {
	// 	chapter := Chapter{}

	// 	fmt.Println(e.DOM)

	// 	fmt.Println(e.Text)

	// 	fmt.Println(e.Name)

	// 	// Получаем заголовок главы
	// 	title := strings.TrimSpace(e.ChildText("span.chapter-title"))
	// 	fmt.Println("Title found:", title)
	// 	chapter.Title = title

	// 	// Получаем дату публикации главы
	// 	date := strings.TrimSpace(e.ChildText("span.text-surface-400.text-right.whitespace-nowrap.text-sm"))
	// 	fmt.Println("Date found:", date)
	// 	chapter.Date = date

	// 	// Получаем ссылку на главу
	// 	link := e.ChildAttr("a", "href")
	// 	fmt.Println("Link found:", link)
	// 	chapter.Link = link

	// 	fmt.Println("Chapter found:", chapter) // Отладочная информация

	// 	manga.Chapters = append(manga.Chapters, chapter)
	// })

	// Начало сбора данных
	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при посещении страницы: %w", err)
	}

	return manga, nil
}

func PrintManga(manga *Manga) {
	if manga == nil {
		fmt.Println("Информация о манге не найдена.")
		return
	}

	fmt.Println("Название:")
	for _, title := range manga.Title {
		fmt.Println("'", title, "'")
	}
	fmt.Println("Описание:", manga.Description)
	fmt.Println("Жанры:")
	for _, genre := range manga.Genre {
		fmt.Println("--", genre)
	}

	fmt.Println("Главы:")
	for _, chapter := range manga.Chapters {
		fmt.Printf("  Глава: %s (%s) - %s\n", chapter.Title, chapter.Link, chapter.Date)
	}
}
