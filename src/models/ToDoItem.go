package models

import (
	"time"

	"github.com/google/uuid"
)

type ToDoItem struct {
	Description string
	DueDate     time.Time
	Id          string
}

func NewToDoItem(description string) *ToDoItem {
	t := ToDoItem{Description: description, DueDate: time.Now().AddDate(0, 0, 7), Id: uuid.New().String()}
	return &t
}
