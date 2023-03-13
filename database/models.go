package database

type User struct {
	Userid   string
	Username string
	Password string
}

type Post struct {
	Id    int64
	Title string
	Body  string
	Time  string
	User  User
}
