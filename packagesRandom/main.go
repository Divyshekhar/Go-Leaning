package main

import (
	"fmt"
	"packagesRandom/user"
	"packagesRandom/auth"
)

func main() {
	auth.LoginWithCred("name", "asod")

	session := auth.GetSession()

	fmt.Println(session)

	user := user.User{
		Email : "johndoe@gmail.com",
		Name: "John doe",
	}
	fmt.Println(user)
}
