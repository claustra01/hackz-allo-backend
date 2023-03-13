package database

type User struct {
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
	Status  bool
	Message string
}
