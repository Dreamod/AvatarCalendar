package media

import (
	"aiCalendarApp/pkg/image"
	"context"
	"errors"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Service interface {
	CreateImage(ctx context.Context, input CreateMediaInput, file io.Reader) (*CreateMediaOutput, error)
	CreateVideo(ctx context.Context, input CreateMediaInput, file io.Reader) (*CreateMediaOutput, error)
}

var (
	ErrCreateUploadSrcDir         = errors.New("create_upload_src_dir")
	ErrCreateUploadPreviewDir     = errors.New("create_upload_preview_dir")
	ErrSaveToDiskFailed           = errors.New("save_to_disk_failed")
	ErrMediaCreateFailed          = errors.New("media_create_failed")
	ErrImageCheckDimensionsFailed = errors.New("image_check_dimensions_failed")
	ErrImageCreatePreviewFailed   = errors.New("image_create_preview_failed")
)

type CreateMediaInput struct {
	Date     time.Time
	FileName string
	FileExt  string
}

type CreateMediaOutput struct {
	ID int
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

// CreateImage логика создания изображения - сохраняет исходник, создает превью, записывает данные в БД
func (service *service) CreateImage(ctx context.Context, input CreateMediaInput, file io.Reader) (*CreateMediaOutput, error) {
	const mediaType = "image"

	// запишем дату в переменную
	date := input.Date

	// определяем папку загрузки исходника
	uploadSrcDir := MakeUploadDirname(date, mediaType, "full")

	// создаем промежуточные директории для исходника
	if err := os.MkdirAll(uploadSrcDir, 0o755); err != nil {
		return nil, ErrCreateUploadSrcDir
	}

	// определяем имя файла
	fileName := MakeFileName(date, input.FileExt)

	// определяем полный путь сохранения исходного файла
	fullSrcPath := filepath.Join(uploadSrcDir, fileName)

	// сохраняем файл
	_, err := saveToDisk(file, fullSrcPath)
	if err != nil {
		return nil, ErrSaveToDiskFailed
	}

	// определяем папку загрузки превью
	uploadPreviewDir := MakeUploadDirname(date, mediaType, "preview")

	// создаем промежуточные директории для превью
	if err := os.MkdirAll(uploadPreviewDir, 0o755); err != nil {
		return nil, ErrCreateUploadPreviewDir
	}

	// определяем полный путь сохранения превью файла
	fullPreviewPath := filepath.Join(uploadPreviewDir, fileName)

	// определяем ориентацию изображения
	isHorizontal, err := image.IsHorizontal(fullSrcPath)
	if err != nil {
		return nil, ErrImageCheckDimensionsFailed
	}

	// определяем разрешение ресайза
	rw := 0
	hr := 0
	if isHorizontal {
		rw = 300
	} else {
		hr = 300
	}

	// создаем превью
	err = image.ResizeImage(fullSrcPath, fullPreviewPath, rw, hr)
	if err != nil {
		return nil, ErrImageCreatePreviewFailed
	}

	// запись данных медиа в БД
	calendarMedia := NewMedia(date, fullSrcPath, fullPreviewPath, mediaType)
	dayId := calendarMedia.DayId

	// если запись существует, то будем ее обновлять
	existedMedia, _ := service.repo.GetByDayId2Type(ctx, dayId, mediaType)
	var resultMediaId int
	if existedMedia == nil {
		// запись в БД
		resultCalendarMedia, err := service.repo.Create(ctx, calendarMedia)
		if err != nil {
			return nil, ErrMediaCreateFailed
		}
		resultMediaId = resultCalendarMedia.Id
	} else {
		resultMediaId = existedMedia.Id
	}

	return &CreateMediaOutput{resultMediaId}, nil
}

// CreateVideo логика создания видео - сохраняет исходник, создает превью, записывает данные в БД
func (service *service) CreateVideo(ctx context.Context, input CreateMediaInput, file io.Reader) (*CreateMediaOutput, error) {
	const mediaType = "video"

	// запишем дату в переменную
	date := input.Date

	// определяем папку загрузки исходника
	uploadSrcDir := MakeUploadDirname(date, mediaType, "full")

	// создаем промежуточные директории для исходника
	if err := os.MkdirAll(uploadSrcDir, 0o755); err != nil {
		return nil, ErrCreateUploadSrcDir
	}

	// определяем имя файла
	fileName := MakeFileName(date, input.FileExt)

	// определяем полный путь сохранения исходного файла
	fullSrcPath := filepath.Join(uploadSrcDir, fileName)

	// сохраняем файл
	_, err := saveToDisk(file, fullSrcPath)
	if err != nil {
		return nil, ErrSaveToDiskFailed
	}

	// определяем папку загрузки превью
	uploadPreviewDir := MakeUploadDirname(date, mediaType, "preview")

	// создаем промежуточные директории для превью
	if err := os.MkdirAll(uploadPreviewDir, 0o755); err != nil {
		return nil, ErrCreateUploadPreviewDir
	}

	// определяем полный путь сохранения превью файла
	fullPreviewPath := filepath.Join(uploadPreviewDir, fileName)

	//TODO convert video to webp

	// запись данных медиа в БД
	calendarMedia := NewMedia(date, fullSrcPath, fullPreviewPath, mediaType)
	dayId := calendarMedia.DayId

	// если запись существует, то будем ее обновлять
	existedMedia, _ := service.repo.GetByDayId2Type(ctx, dayId, mediaType)
	var resultMediaId int
	if existedMedia == nil {
		// запись в БД
		resultCalendarMedia, err := service.repo.Create(ctx, calendarMedia)
		if err != nil {
			return nil, ErrMediaCreateFailed
		}
		resultMediaId = resultCalendarMedia.Id
	} else {
		resultMediaId = existedMedia.Id
	}

	return &CreateMediaOutput{resultMediaId}, nil
}
