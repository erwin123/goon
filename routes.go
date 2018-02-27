package main

import (
	"net/http"

	handlers "github.com/goon/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", handlers.Index},
	Route{"TodoIndex", "GET", "/todos", handlers.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", handlers.TodoShow},
	Route{"TodoCreate", "POST", "/todos", handlers.TodoCreate},
}
