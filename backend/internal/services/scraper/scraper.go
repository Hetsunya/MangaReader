package scraper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	models "manga-reader/backend/internal/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// Scrap - функция, которая парсит веб-страницу и получает информацию о манге, а затем возвращает структуру Manga.
//
// Она принимает строку URL в качестве параметра и возвращает указатель на структуру Manga и ошибку.
// Функция использует библиотеку Colly для парсинга веб-страницы и собирает информацию о названии манги, жанрах, описании и главах.
// Она устанавливает ограничения для запросов и обрабатывает разные элементы HTML, чтобы заполнить структуру Manga.
// Если при посещении веб-страницы возникла ошибка, то функция возвращает nil и ошибку.
func Scrap(url string) (*models.Manga, error) {
	c := colly.NewCollector()

	// Установка лимитов для запросов
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,               // Запросов в секунду
		Delay:       1 * time.Second, // Можно уменьшить, если допустимо
	})

	manga := &models.Manga{}

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

	// Получение количества глав и статуса манги
	c.OnHTML("span", func(e *colly.HTMLElement) {
		text := e.Text
		if strings.Contains(text, "Глав:") {
			manga.NumberOfChapters = strings.TrimSpace(strings.Replace(text, "Глав:", "", 1))
		}
		if strings.Contains(text, "Статус:") {
			manga.Status = strings.TrimSpace(strings.Replace(text, "Статус:", "", 1))
		}
	})

	// Получение года выхода манга
	c.OnHTML(".badge.variant-soft-tertiary", func(e *colly.HTMLElement) {
		manga.Year = e.Text
	})

	// Начало сбора данных
	err := c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка при посещении страницы: %w", err)
	}

	return manga, nil
}

/*
 */
func ScrapChapters(query string, manga *models.Manga, baseURL string) error {
	page := 1

	for {
		url := fmt.Sprintf("%s&page=%d&sort=desc", query, page)

		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("Ошибка при загрузке страницы: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("Ошибка при загрузке страницы: статус %d", resp.StatusCode)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return fmt.Errorf("Ошибка при создании документа goquery: %v", err)
		}

		foundChapters := 0

		doc.Find("li.flex.justify-between.card.variant-soft-surface.mb-3.p-2.md\\:px-3").Each(func(i int, s *goquery.Selection) {
			chapterTitle := s.Find("span.chapter-title").Text()
			chapterLink, _ := s.Find("a").Attr("href")
			publicationDate := s.Find("span.text-surface-400.text-right.whitespace-nowrap.text-sm").Text()

			chapter := models.Chapter{
				Title: strings.TrimSpace(chapterTitle),
				Link:  baseURL + strings.TrimSpace(chapterLink),
				Date:  strings.TrimSpace(publicationDate),
			}

			manga.Chapters = append(manga.Chapters, chapter)
			foundChapters++
		})

		// Если больше нет глав, выходим из цикла
		if foundChapters == 0 {
			break
		}

		page++
	}

	return nil
}

// PrintManga выводит информацию о Манге.
//
// Принимает указатель на объект Манга в качестве параметра и выводит название,
// описание, жанры и главы Манги. Если объект Манга равен nil, выводится сообщение об ошибке.
func PrintManga(manga *models.Manga) {
	if manga == nil {
		fmt.Println("Информация о манге не найдена.")
		return
	}

	fmt.Println("Название:")
	for _, title := range manga.Title {
		fmt.Println("'", title, "'")
	}
	fmt.Println("Количество глав:", manga.NumberOfChapters)
	fmt.Println("Статус:", manga.Status)
	fmt.Println("Описание:", manga.Description)
	fmt.Println("Год:", manga.Year)
	fmt.Println("Жанры:")
	for _, genre := range manga.Genre {
		fmt.Println("--", genre)
	}

	fmt.Println("Главы:")
	for _, chapter := range manga.Chapters {
		fmt.Printf("  Глава: %s (%s) - %s\n", chapter.Title, chapter.Link, chapter.Date)
	}
}
