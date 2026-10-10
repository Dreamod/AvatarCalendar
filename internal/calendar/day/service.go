package day

import (
	"context"
	"errors"
	"time"
)

var (
	ErrDayAlreadyExists = errors.New("day_already_exists")
	ErrInvalidDate      = errors.New("invalid_date_format")
)

type Service interface {
	CreateDay(ctx context.Context, req *CreateRequest) (*Day, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (service *service) CreateDay(ctx context.Context, req *CreateRequest) (*Day, error) {
	// валидация/парсинг даты
	date, err := time.Parse("02.01.2006", req.Date)
	if err != nil {
		return nil, ErrInvalidDate
	}

	// создаём модель
	calendarDay := NewDay(date, req.Title, req.Description)
	dayID := calendarDay.Id

	// проверяем, нет ли на этот день записи
	existedDay, _ := service.repo.GetById(ctx, dayID)
	if existedDay != nil {
		return nil, ErrDayAlreadyExists
	}

	// создаем запись
	createdDay, err := service.repo.Create(ctx, calendarDay)
	if err != nil {
		return nil, err
	}

	return createdDay, nil
}
