package media

import (
	"aiCalendarApp/configs"
	fileValidator "aiCalendarApp/pkg/file"
	"aiCalendarApp/pkg/response"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	*configs.Config
	Service Service
}

type HandlerDeps struct {
	*configs.Config
	Service Service
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:  deps.Config,
		Service: deps.Service,
	}
	router.HandleFunc("POST /api/media/image/{date}", handler.CreateImage())
	router.HandleFunc("POST /api/media/video/{date}", handler.CreateVideo())
	//router.HandleFunc("GET /{date}", handler.GoTo())
	//router.HandleFunc("PATCH /calendar/{date}", handler.Update())
}

// CreateImage сохраняет картинку на диск, создает превью и запись в БД
func (handler *Handler) CreateImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		const maxSize = 20 * 1024 * 1024 // 20 МБ

		// подготовка ответа
		output := CreateResponse{
			Success: false,
			Data:    "",
			Errors:  make([]string, 0),
		}

		// получаем дату
		dateStr := r.PathValue("date")
		date, err := time.Parse("02.01.2006", dateStr)
		if err != nil {
			output.Errors = append(output.Errors, "неверный формат даты. Ожидается dd.mm.YYYY")
			response.Json(w, output, http.StatusBadRequest)
			return
		}

		// проверка наличия файла
		file, header, err := r.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				output.Errors = append(output.Errors, "file: не заполнено обязательное поле")
				response.Json(w, output, http.StatusBadRequest)
				return
			}
			// другие ошибки FormFile (повреждённый multipart и т.п.)
			output.Errors = append(output.Errors, "file: некорректные данные")
			response.Json(w, output, http.StatusBadRequest)
			return
		}
		defer file.Close()

		// получим расширение файла
		ext := strings.ToLower(filepath.Ext(header.Filename))

		// проверка "веса" файла
		if header.Size > maxSize {
			output.Errors = append(output.Errors, "file: файл слишком большой > ", strconv.Itoa(maxSize), "MB")
			response.Json(w, output, http.StatusBadRequest)
		}

		// валидация файла картинки
		err = fileValidator.IsValidImage(file, ext)
		if err != nil {
			output.Errors = append(output.Errors, err.Error())
			response.Json(w, output, http.StatusBadRequest)
		}

		// возвращаем указатель чтения в начало
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			output.Errors = append(output.Errors, "file: ошибка чтения файла")
			response.Json(w, output, http.StatusInternalServerError)
			return
		}

		// выполняем бизнес логику
		CreteMediaInput := CreateMediaInput{
			Date:     date,
			FileName: header.Filename,
			FileExt:  ext,
		}
		result, err := handler.Service.CreateImage(ctx, CreteMediaInput, file)
		if err != nil {
			output.Errors = append(output.Errors, "ошибка при сохранении файла")
			response.Json(w, output, http.StatusInternalServerError)
			return
		}

		// формируем успешный ответ
		output = CreateResponse{
			Success: true,
			Data: map[string]string{
				"id": strconv.Itoa(result.ID),
			},
			Errors: nil,
		}
		response.Json(w, output, http.StatusCreated)
	}
}

// CreateVideo сохраняет видео на диск, создает превью и запись в БД
func (handler *Handler) CreateVideo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		const maxSize = 20 * 1024 * 1024 // 20 МБ

		// подготовка ответа
		output := CreateResponse{
			Success: false,
			Data:    "",
			Errors:  make([]string, 0),
		}

		// получаем дату
		dateStr := r.PathValue("date")
		date, err := time.Parse("02.01.2006", dateStr)
		if err != nil {
			output.Errors = append(output.Errors, "неверный формат даты. Ожидается dd.mm.YYYY")
			response.Json(w, output, http.StatusBadRequest)
			return
		}

		// проверка наличия файла
		file, header, err := r.FormFile("file")
		if err != nil {
			if errors.Is(err, http.ErrMissingFile) {
				output.Errors = append(output.Errors, "file: не заполнено обязательное поле")
				response.Json(w, output, http.StatusBadRequest)
				return
			}
			// другие ошибки FormFile (повреждённый multipart и т.п.)
			output.Errors = append(output.Errors, "file: некорректные данные")
			response.Json(w, output, http.StatusBadRequest)
			return
		}
		defer file.Close()

		// получим расширение файла
		ext := strings.ToLower(filepath.Ext(header.Filename))

		// проверка "веса" файла
		if header.Size > maxSize {
			output.Errors = append(output.Errors, "file: файл слишком большой > ", strconv.Itoa(maxSize), "MB")
			response.Json(w, output, http.StatusBadRequest)
		}

		// валидация файла видео
		err = fileValidator.IsValidVideo(file, ext)
		if err != nil {
			output.Errors = append(output.Errors, err.Error())
			response.Json(w, output, http.StatusBadRequest)
		}

		// возвращаем указатель чтения в начало
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			output.Errors = append(output.Errors, "file: ошибка чтения файла")
			response.Json(w, output, http.StatusInternalServerError)
			return
		}

		// выполняем бизнес логику
		CreteMediaInput := CreateMediaInput{
			Date:     date,
			FileName: header.Filename,
			FileExt:  ext,
		}
		result, err := handler.Service.CreateVideo(ctx, CreteMediaInput, file)
		if err != nil {
			output.Errors = append(output.Errors, "ошибка при сохранении файла")
			response.Json(w, output, http.StatusInternalServerError)
			return
		}

		// формируем успешный ответ
		output = CreateResponse{
			Success: true,
			Data: map[string]string{
				"id": strconv.Itoa(result.ID),
			},
			Errors: nil,
		}
		response.Json(w, output, http.StatusCreated)
	}
}
