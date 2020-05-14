package request

type CreateRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Status Status `json:"status"`
	Tag string `json:"tag"`
}

type UpdateRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Status Status `json:"status"`
	Tag string `json:"tag"`
}

type Status int
const (
	OnHold Status = 0
	InProcess Status = 1
	Done Status = 2
)