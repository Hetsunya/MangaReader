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

// SearchResult представляет результат поиска манги
type SearchResult struct {
	FoundMangas []FoundManga
}

// FoundManga представляет информацию о найденной манге
type FoundManga struct {
	URL   string
	Title string
}

//Все ниже для БД
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type MangaList struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	URL    string `json:"url"`
	Status string `json:"status"` // например, 'читаю', 'в планах', 'брошено'
}

type MangaTag struct {
	ID     int    `json:"id"`
	ListID int    `json:"list_id"`
	Tag    string `json:"tag"`
}
