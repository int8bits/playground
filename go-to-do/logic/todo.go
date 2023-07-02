package logic

import (
	"go-to-do/models"
	"time"
)

func NewToDo(id int, name, description, notes string) *models.ToDo {
	return &models.ToDo{
		Id:              id,
		Name:            name,
		Description:     description,
		AdditionalNotes: notes,
		CreatedAt:       time.Now(),
	}
}
