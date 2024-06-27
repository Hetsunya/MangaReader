package models

// Manga представляет информацию о манге
type Manga struct {
	Title            []string
	Description      string
	Genre            []string
	Chapters         []Chapter
	NumberOfChapters string
	Status           string
	Year             string
}

// Chapter представляет информацию о главе манги
type Chapter struct {
	Title string
	Date  string
	Link  string
	Pages MangaPage
}

type MangaPage struct {
	ImageURL string  // URL изображения
	Chapter  Chapter // Глава манги, к которой относится страница
	// Дополнительные поля могут быть добавлены, например:
	// PageNumber int    // Номер страницы
	// ChapterTitle string // Заголовок главы
}
