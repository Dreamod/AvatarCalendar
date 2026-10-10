package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

// IsHorizontal определяет ориентацию изображения
func IsHorizontal(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	cfg, _, err := image.DecodeConfig(file)
	if err != nil {
		return false, err
	}

	return cfg.Width > cfg.Height, nil
}

// ResizeImage ресайз изображений
func ResizeImage(inputPath string, outputPath string, width int, height int) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	var srcImg image.Image
	ext := strings.ToLower(filepath.Ext(inputPath))

	// декодирование по формату
	switch ext {
	case ".jpg", ".jpeg":
		srcImg, err = jpeg.Decode(file)
	case ".png":
		srcImg, err = png.Decode(file)
	default:
		return fmt.Errorf("поддерживаются только JPEG и PNG, получен формат: %s", ext)
	}
	if err != nil {
		return fmt.Errorf("ошибка декодирования изображения: %w", err)
	}

	// ресайз: если один из размеров 0 — пропорции сохраняются
	dstImg := resize.Resize(uint(width), uint(height), srcImg, resize.Lanczos3)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("не удалось создать выходной файл: %w", err)
	}
	defer outFile.Close()

	// кодирование в нужный формат
	switch ext {
	case ".jpg", ".jpeg":
		opts := &jpeg.Options{Quality: 85} // 0–100, баланс качества и размера
		return jpeg.Encode(outFile, dstImg, opts)
	case ".png":
		return png.Encode(outFile, dstImg)
	}

	return nil
}
