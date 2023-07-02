package models

import "time"

type ToDo struct {
	Id              int
	Name            string
	Description     string
	IsActive        bool
	Status          int
	CreatedAt       time.Time
	UpdateAt        time.Time
	FinisedAt       time.Time
	AdditionalNotes string
	Flow            []int
}
