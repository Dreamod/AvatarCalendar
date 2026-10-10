package day

import (
	"aiCalendarApp/configs"
	"aiCalendarApp/pkg/request"
	"aiCalendarApp/pkg/response"
	"net/http"
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
	router.HandleFunc("POST /api/day", handler.Create())
	//router.HandleFunc("GET /{date}", handler.GoTo())
	//router.HandleFunc("PATCH /calendar/{date}", handler.Update())
}

func (handler *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := request.HandleBody[CreateRequest](&w, r)
		if err != nil {
			// ошибка декодирования — сразу отвечаем
			response.Json(w, CreateResponse{
				Success: false,
				Errors:  []string{"ошибка чтения тела запроса"},
			}, http.StatusBadRequest)
			return
		}

		output := CreateResponse{
			Success: false,
			Data:    "",
			Errors:  make([]string, 0),
		}

		// выполняем бизнес логику
		createdDay, err := handler.Service.CreateDay(ctx, body)
		if err != nil {
			switch err {
			case ErrDayAlreadyExists:
				output.Errors = append(output.Errors, "на указанную дату уже есть запись в календаре")
				response.Json(w, output, http.StatusConflict) // 409
				return
			case ErrInvalidDate:
				output.Errors = append(output.Errors, "date: некорректный формат даты")
				response.Json(w, output, http.StatusBadRequest) // 400
				return
			default:
				// любая другая ошибка (в т.ч. БД)
				output.Errors = append(output.Errors, "внутренняя ошибка сервиса")
				response.Json(w, output, http.StatusInternalServerError) // 500
				return
			}
		}

		// формируем успешный ответ
		output = CreateResponse{
			Success: true,
			Data: map[string]string{
				"id": createdDay.Id,
			},
			Errors: nil,
		}
		response.Json(w, output, http.StatusCreated)
	}
}
