package database

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	UserId   string
	Name     string
	Password string
}

type Post struct {
	Id     uuid.UUID
	Title  string
	Body   string
	Time   string
	UserId string
}

type Response struct {
	Result  string
	Message string
}
