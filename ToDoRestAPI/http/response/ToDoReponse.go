package response

import (
	"time"
)

type ToDosResponse struct {
	ToDos []*ToDo `json:"todolist"`
}

type ToDo struct {
	Id int `json:"id"`
	UserId int `json:"userid"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status Status `json:"status"`
	Tag string `json:"tag"`
	CreatedDate time.Time `json:time`
}

type Status int
const (
	OnHold Status = 0
	InProcess Status = 1
	Done Status = 2
)