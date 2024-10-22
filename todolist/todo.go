package main

type Todo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	IsComplete bool   `json:"is_complete"`
}
