package main

import (
	"localhost/greetings/entities/user"
	"localhost/greetings/notification"
)

func main() {

	user := user.NewUser(user.UserProps{Email: "fulano@email.com"})

	println(notification.GetErrorMessages())
	println(user.Id)
	println(user.Password)
	println(user.Email)
}
