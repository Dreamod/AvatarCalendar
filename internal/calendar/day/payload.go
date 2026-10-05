package day

type CreateRequest struct {
	Date        string `json:"date" validate:"required,datetime=02.01.2006"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateResponse struct {
	Success bool     `json:"success"`
	Data    any      `json:"data"`
	Errors  []string `json:"errors"`
}
