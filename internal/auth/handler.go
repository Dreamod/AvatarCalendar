package auth

import (
	"aiCalendarApp/configs"
	"aiCalendarApp/pkg/request"
	"aiCalendarApp/pkg/response"
	"fmt"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		data := LoginResponse{
			Token: "1234567890",
		}
		fmt.Println(payload)
		response.Json(w, data, http.StatusOK)
	}
}
