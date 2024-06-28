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
	Pages MangaPage // Пока хз, мейби удалить стоит
}

type MangaPage struct {
	ImageURL string // URL изображения
	// Дополнительные поля могут быть добавлены, например:
	// PageNumber int    // Номер страницы
	// ChapterTitle string // Заголовок главы
}

const (
	BaseURL      = "https://mangapoisk.live"
	ChapterParse = "?tab=chapters"
	//TODO: Разобраться что делать с путем
	JsonDir = "C:\\Users\\vital\\Desktop\\MangaReader\\backend\\jsons"
)
