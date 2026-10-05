package day

import (
	"aiCalendarApp/configs"
	"aiCalendarApp/pkg/request"
	"aiCalendarApp/pkg/response"
	"net/http"
	"time"
)

type DayHandler struct {
	*configs.Config
	DayRepository *DayRepository
}

type DayHandlerDeps struct {
	*configs.Config
	DayRepository *DayRepository
}

func NewDayHandler(router *http.ServeMux, deps DayHandlerDeps) {
	handler := &DayHandler{
		Config:        deps.Config,
		DayRepository: deps.DayRepository,
	}
	router.HandleFunc("POST /calendar", handler.Create())
	//router.HandleFunc("GET /{date}", handler.GoTo())
	//router.HandleFunc("PATCH /calendar/{date}", handler.Update())
}

func (handler *DayHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получим данные из тела запроса
		body, err := request.HandleBody[CreateRequest](&w, r)
		if err != nil {
			return
		}
		// подготовка ответа
		output := CreateResponse{
			Success: false,
			Data:    "",
			Errors:  make([]string, 0),
		}
		// конвертируем строку в дату
		date, err := time.Parse("02.01.2006", body.Date)
		if err != nil {
			output.Errors = append(output.Errors, "date: некорректный формат даты")
			response.Json(w, output, http.StatusBadRequest)
			return
		}
		// создаем день
		calendarDay := NewDay(date, body.Title, body.Description)
		dayId := calendarDay.Id
		// существует ли день
		existedDay, _ := handler.DayRepository.GetById(dayId)
		if existedDay != nil {
			output.Errors = append(output.Errors, "на указанную дату уже есть запись в календаре")
			response.Json(w, output, http.StatusBadRequest)
			return
		}
		// создадим день
		createdCalendarDay, err := handler.DayRepository.Create(calendarDay)
		if err != nil {
			output.Errors = append(output.Errors, "при создании записи дня произошла ошибка")
			response.Json(w, output, http.StatusBadRequest)
			return
		}
		//пишем успешный ответ
		output = CreateResponse{
			Success: true,
			Data: map[string]string{
				"id": createdCalendarDay.Id,
			},
			Errors: nil,
		}
		response.Json(w, output, http.StatusCreated)
	}
}
