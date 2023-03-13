package database

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	UserId   string
	Name     string
	Password string
}

type Post struct {
	Id    int64
	Title string
	Body  string
	Time  string
	User  User
}

type Response struct {
	Result  string
	Message string
}
