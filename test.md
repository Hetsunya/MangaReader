Отлично, вот пример документации в формате Markdown:

# Документация для MangaReader

## Общие сведения

**Проект**: MangaReader

**Описание**: Приложение для управления списками манги с различными статусами.

**Технологии**: Go, SQLite, gomobile

## Установка и настройка

### Клонирование репозитория

```bash
git clone https://github.com/yourusername/MangaReader.git
cd MangaReader
```

### Настройка окружения

1. Установите Go.
2. Установите gomobile:
    ```bash
    go install golang.org/x/mobile/cmd/gomobile@latest
    gomobile init
    ```
3. Установите SQLite3.
4. Убедитесь, что у вас установлен gcc.

### Запуск

Для запуска приложения используйте команду:

```bash
go run backend/cmd/main.go
```

## Структура проекта

```plaintext
MangaReader
│
├── backend
│   ├── cmd
│   │   └── main.go
│   │
│   ├── internal
│   │   ├── db
│   │   │   ├── db.go
│   │   │   └── manga_lists.go
│   │   │
│   │   ├── lib
│   │   │   └── jsonutil
│   │   │       └── jsonutil.go
│   │   │
│   │   ├── models
│   │   │   └── models.go
│   │   │
│   │   ├── services
│   │   │   ├── imageextractor
│   │   │   │   └── imageextractor.go
│   │   │   │
│   │   │   ├── scraper
│   │   │   │   └── scraper.go
│   │   │   │
│   │   │   └── searcher
│   │   │       └── searcher.go
│   │   │
│   │   └── utils
│   │       └── utils.go
│   │
│   └── go.mod
│
└── README.md
```

## Структуры данных

### MangaList

```go
type MangaList struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    URL    string `json:"url"`
    Status string `json:"status"`
}
```

### MangaTag

```go
type MangaTag struct {
    ID     int    `json:"id"`
    ListID int    `json:"list_id"`
    Tag    string `json:"tag"`
}
```

## API документация

### Методы работы с БД

#### Создание списка манги

**Функция**: `CreateMangaList`

**Описание**: Добавляет новый список манги в базу данных.

**Пример**:

```go
mangaList := models.MangaList{
    Name:   "повышение уровня",
    URL:    selectedMangaURL,
    Status: "читаю",
}

listID, err := db.CreateMangaList(database, mangaList)
if err != nil {
    log.Error("Error creating manga list:", err)
}
fmt.Println("Created manga list with ID:", listID)
```

#### Получение всех списков манги

**Функция**: `GetAllMangaLists`

**Описание**: Возвращает все списки манги из базы данных.

**Пример**:

```go
lists, err := db.GetAllMangaLists(database)
if err != nil {
    log.Error("Error fetching manga lists:", err)
}
fmt.Println("Manga lists:", lists)
```

### Примеры использования

- **Добавление манги в список**:
  ```go
  mangaList := models.MangaList{
      Name:   "повышение уровня",
      URL:    "https://example.com/manga/1",
      Status: "читаю",
  }

  listID, err := db.CreateMangaList(database, mangaList)
  if err != nil {
      log.Error("Error creating manga list:", err)
  }
  fmt.Println("Created manga list with ID:", listID)
  ```

- **Получение всех списков манги**:
  ```go
  lists, err := db.GetAllMangaLists(database)
  if err != nil {
      log.Error("Error fetching manga lists:", err)
  }
  fmt.Println("Manga lists:", lists)
  ```
---