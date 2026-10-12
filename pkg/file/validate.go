package file

import (
	"errors"
	"io"
	"net/http"
)

var (
	ErrMissingFile         = errors.New("file: не заполнено обязательное поле")
	ErrUnsupportedImageExt = errors.New("file: несовместимый формат, ожидается .jpg, .jpeg или .png")
	ErrUnsupportedVideoExt = errors.New("file: несовместимый формат, ожидается .mp4")
	ErrUnsupportedMIME     = errors.New("file: передан неподдерживаемый тип файла")
)

func IsValidImage(file io.Reader, ext string) error {

	// проверка типов/форматов файла
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return ErrUnsupportedImageExt
	}

	// проверка MIME по первым байтам
	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	if n == 0 {
		return ErrMissingFile
	}
	mimeType := http.DetectContentType(buf)
	if mimeType != "image/jpeg" && mimeType != "image/png" {
		return ErrUnsupportedMIME
	}

	return nil
}

func IsValidVideo(file io.Reader, ext string) error {

	// проверка типов/форматов файла
	if ext != ".mp4" {
		return ErrUnsupportedVideoExt
	}

	// проверка MIME по первым байтам
	buf := make([]byte, 512)
	n, _ := file.Read(buf)
	if n == 0 {
		return ErrMissingFile
	}
	mimeType := http.DetectContentType(buf)
	if mimeType != "video/mp4" {
		return ErrUnsupportedMIME
	}

	return nil
}
