package main

import "fmt"

const (
	ADMIN = "admin"
)

func main()  {
	user := getUser("Denis","admin",29)
	fmt.Println("Is Admin", user.IsAdmin())
	fmt.Println("CheckAge", user.checkAge())
}

type User struct {
	Name string
	Role string
	Age int
}

type UserInterface interface {
	IsAdmin() bool
	checkAge() bool
}

func (u *User) IsAdmin() bool  {
	return u.Role == "admin"
}

func (u *User) checkAge() bool  {
	return u.Age > 18
}

func getUser(Name string, Role string, Age int) UserInterface  {
	return &User{Name: Name,Role: Role, Age : Age}
}
