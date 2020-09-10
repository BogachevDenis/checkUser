package main

import "fmt"

type User struct {
	Id   int
	Name string
	Role string
	Age int
}


func (u *User) CreateUser(id int, name string, role string, age int) {
	u.Id = id
	u.Name = name
	u.Role = role
	u.Age = age
}

func (u *User) checkAge() bool  {
	return u.Age > 18
}


func main() {
	usr := new(User)
	usr.CreateUser(5, "Denis", "Admin", 24)
	fmt.Println(usr)
	fmt.Println("checkAge - ", usr.checkAge())
}
