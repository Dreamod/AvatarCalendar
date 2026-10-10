package media

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// saveFileToDisk сохраняет файл из io.Reader в папку assets/image/full/
// Возвращает относительный путь к файлу (для сохранения в БД) и ошибку.
func saveToDisk(file io.Reader, filePath string) (string, error) {
	// создаём файл для записи
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Ошибка создания файла %s: %v", filePath, err)
		return "", err
	}
	defer out.Close() // Обязательно закрываем файл при выходе
	// копируем данные из входного потока (файла из запроса) в файл на диске
	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Ошибка копирования данных в %s: %v", filePath, err)
		// Файл может остаться пустым, ОС его создаст, но данных не будет.
		// Можно добавить логику удаления пустого файла, если нужно.
		return "", err
	}
	// возвращаем относительный путь (без точки в начале), чтобы положить в БД
	return filePath, nil
}

func MakeUploadDirname(date time.Time, mediaType string, fileType string) string {
	baseDir := "."
	return filepath.Join(
		baseDir,
		"assets",
		mediaType,
		fileType,
		date.Format("2006"),
		date.Format("01"),
	)
}

func MakeFileName(date time.Time, extension string) string {
	return date.Format("02.01.2006") + extension
}
