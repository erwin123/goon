package repo

import (
	"fmt"

	m "github.com/goon/models"
)

var currentId int
var Todos m.Todos

// Give us some seed data
func Init() {

}

func RepoFindTodo(id int) m.Todo {
	for _, t := range Todos {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return m.Todo{}
}

func RepoCreateTodo(t m.Todo) m.Todo {
	currentId += 1
	t.ID = currentId
	Todos = append(Todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
	for i, t := range Todos {
		if t.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
