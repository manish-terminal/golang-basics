package main

import (
	"fmt"
	"somethingevenidk/auth"
	"somethingevenidk/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("manish", "password")
	session := auth.GetSession()
	fmt.Println(session)
	user := user.User{Email: "email@mail.com", Name: "John Doe"}
	fmt.Println(user)
	color.Red("hello")

}
