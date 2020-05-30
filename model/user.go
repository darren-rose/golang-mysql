package model

type User struct {
	Id       int
	Email    string
	Password string
	Name     string
}

type Tenancy struct {
	Id         int
	UserId     int
	PropertyId int
	StartDate  string
	EndDate    string
	Address    string
}
