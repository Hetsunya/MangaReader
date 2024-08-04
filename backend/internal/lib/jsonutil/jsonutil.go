package jsonutil

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manga-reader/backend/internal/models"
	"os"
	"path/filepath"
	"time"
)

// ToJSON принимает значение и возвращает его строковое представление в формате JSON.
// Кроме того, функция сохраняет JSON-файл по указанному пути, если файл с такими же данными еще не существует.
func ToJSON(v interface{}) (string, error) {
	// Преобразование значения в формат JSON с отступами.
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return "", err
	}

	// Создание хэша данных.
	hash := sha256.Sum256(data)
	hashString := hex.EncodeToString(hash[:])

	// Создание каталога, если он не существует.
	err = os.MkdirAll(models.JsonDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Проверка, существует ли файл с таким же хэшем.
	files, err := ioutil.ReadDir(models.JsonDir)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			existingData, err := ioutil.ReadFile(filepath.Join(models.JsonDir, file.Name()))
			if err != nil {
				return "", err
			}
			existingHash := sha256.Sum256(existingData)
			existingHashString := hex.EncodeToString(existingHash[:])
			if hashString == existingHashString {
				return string(data), nil // Файл уже существует, возвращаем данные без сохранения.
			}
		}
	}

	// Определение уникального имени файла.
	timestamp := time.Now().Format("20060102_150405")
	outputFile := filepath.Join(models.JsonDir, fmt.Sprintf("output_%s.json", timestamp))

	// Создание (или перезапись) JSON-файла.
	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		return "", err
	}

	// Возвращение строкового представления JSON.
	return string(data), nil
}

//TODO: Удалить и сделать на фронте или хз че там
