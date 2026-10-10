package media

type CreateResponse struct {
	Success bool     `json:"success"`
	Data    any      `json:"data"`
	Errors  []string `json:"errors"`
}
