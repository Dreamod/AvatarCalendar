package string

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

// GetMD5Hash создает хэш от переданной строки
func GetMD5Hash(string string) string {
	// создаём хеш-объект
	hasher := md5.New()

	// добавляем данные в хеш
	hasher.Write([]byte(string))

	// получаем итоговый хеш-сумму в виде байтового слайса
	hashBytes := hasher.Sum(nil)

	// преобразуем байты в шестнадцатеричную строку
	return hex.EncodeToString(hashBytes)
}

// GetRandomString получает рандомную строку заданной длины
func GetRandomString(n int) string {
	// слайс символов для результата
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// создает слайс символов заданной длины
	b := make([]rune, n)

	// в цикле добавляет в слайс псевдослучайный символ из слайса рун
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
