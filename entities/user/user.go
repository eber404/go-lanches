package user

import (
	"localhost/greetings/notification"
	"regexp"

	"github.com/google/uuid"
)

type UserProps struct {
	Email    string
	Password string
}

type User struct {
	Id       string
	Email    string
	Password string
}

func validateEmail(email string) bool {
	isEmailFormat, _ := regexp.MatchString("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$", email)

	return isEmailFormat
}

func validatePasword(password string) bool {
	hasValidLength := len(password) >= 8

	return hasValidLength
}

func NewUser(props UserProps) User {

	isEmailValid := validateEmail(props.Email)
	isPasswordValid := validatePasword(props.Password)

	if !isEmailValid {
		notification.Add("inform a valid email address", notification.Error)
	}

	if !isPasswordValid {
		notification.Add("passwords should have at least 8 characters", notification.Error)
	}

	user := User{
		Email:    props.Email,
		Password: props.Password,
	}

	user.Id = uuid.NewString()

	return user
}
