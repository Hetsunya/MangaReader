package models

// Manga представляет информацию о манге
// @Description Manga represents information about a manga
// @Accept json
// @Produce json
type Manga struct {
	// @Description Заголовок манги
	// @Example ["One Piece", "Naruto"]
	Title []string `json:"title"`
	// @Description Описание манги
	// @Example "A great manga about pirates."
	Description string `json:"description"`
	// @Description Жанры манги
	// @Example ["Action", "Adventure"]
	Genre []string `json:"genre"`
	// @Description Список глав манги
	Chapters []Chapter `json:"chapters"`
	// @Description Количество глав
	// @Example "150"
	NumberOfChapters string `json:"number_of_chapters"`
	// @Description Статус манги
	// @Example "Ongoing"
	Status string `json:"status"`
	// @Description Год выпуска
	// @Example "2000"
	Year string `json:"year"`
}

// Chapter представляет информацию о главе манги
// @Description Chapter represents information about a manga chapter
// @Accept json
// @Produce json
type Chapter struct {
	// @Description Заголовок главы
	// @Example "Chapter 1: Departure"
	Title string `json:"title"`
	// @Description Дата публикации главы
	// @Example "2024-01-01"
	Date string `json:"date"`
	// @Description Ссылка на главу
	// @Example "https://example.com/chapter1"
	Link string `json:"link"`
	// @Description Страницы главы манги
	// @Example [{"image_url": "https://example.com/image1.jpg"}]
	Pages MangaPage `json:"pages"`
}

type MangaPage struct {
	// @Description URL изображения страницы
	// @Example "https://example.com/page1.jpg"
	ImageURL string `json:"image_url"`
}

const (
	BaseURL      = "https://mangapoisk.live"
	ChapterParse = "?tab=chapters"
)

// SearchResult представляет результат поиска манги
// @Description SearchResult represents the result of a manga search
// @Accept json
// @Produce json
type SearchResult struct {
	// @Description Список найденных манг
	// @Example [{"url": "https://example.com/manga1", "title": "Manga Title 1"}]
	FoundMangas []FoundManga `json:"found_mangas"`
}

// FoundManga представляет информацию о найденной манге
// @Description FoundManga represents information about a found manga
// @Accept json
// @Produce json
type FoundManga struct {
	// @Description URL найденной манги
	// @Example "https://example.com/manga1"
	URL string `json:"url"`
	// @Description Заголовок найденной манги
	// @Example "Manga Title 1"
	Title string `json:"title"`
}

// ErrorResponse представляет ошибку в ответе API
type ErrorResponse struct {
	Error string `json:"error"`
}
