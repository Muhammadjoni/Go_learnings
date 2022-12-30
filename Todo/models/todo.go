package models

type Todo struct {
	Id   uint   `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}
