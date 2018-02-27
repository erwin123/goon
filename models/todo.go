package models

import "time"

//model todo
type Todo struct {
	ID        int       `json:"ID"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

//modellist todo
type Todos []Todo
