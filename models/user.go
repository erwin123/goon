package models

import "time"

//model user
type User struct {
	id           int
	deleted      int
	username     string
	accountno    string
	passwordsalt string
	passwordhash string
	wrongcount   int
	locked       int
	createdat    time.Time
	createdby    string
	updatedat    time.Time
	updatedby    string
}

//modellist user
type Users []User
