package main

type Priority string

const (
	Low Priority = "low"
	Medium Priority = "medium"
	High Priority = "high"
)

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
	Priority Priority `json:"priority"`
}