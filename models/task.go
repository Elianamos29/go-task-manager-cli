package models

import "time"

type Priority string

const (
	Low 	Priority = "low"
	Medium 	Priority = "medium"
	High 	Priority = "high"
)

type Task struct {
	ID 			int `json:"id" gorm:"primaryKey;autoIncrement"`
	Name 		string `json:"name"`
	Done 		bool `json:"done"`
	Priority	Priority `json:"priority"`
	DueDate		time.Time `json:"due_date"`
	Tags		string `json:"tags"`
}