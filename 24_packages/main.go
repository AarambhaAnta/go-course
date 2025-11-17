package main

import (
	"fmt"

	"github.com/AarambhaAnta/podcast/auth"
	"github.com/AarambhaAnta/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("aditya kumar", ".---")

	session := auth.GetSession()
	fmt.Println("session: ", session)

	user := user.User{
		Email: "user@email.com",
		Name:  "John Doe",
	}

	fmt.Println("email: ", user.Email)
	fmt.Println("name: ", user.Name)

	color.Green("true")
	color.Red("false")
}
