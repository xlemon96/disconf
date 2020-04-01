package http

type CreateAppRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
