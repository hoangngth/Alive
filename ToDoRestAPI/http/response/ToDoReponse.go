package response

import (
	"time"
)

type ToDosResponse struct {
	ToDos []*ToDo `json:"todolist"`
}

type ToDo struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	Tag         string    `json:"tag"`
	CreatedDate time.Time `json:time`
}

type Status int

const (
	ONHOLD    Status = 0
	INPROCESS Status = 1
	DONE      Status = 2
)
