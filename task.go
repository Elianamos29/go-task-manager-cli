package main

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Done bool `json:"done"`
}