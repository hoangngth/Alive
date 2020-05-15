package request

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	Tag         string `json:"tag"`
}

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	Tag         string `json:"tag"`
}

type Status int

const (
	ONHOLD    Status = 0
	INPROCESS Status = 1
	DONE      Status = 2
)
