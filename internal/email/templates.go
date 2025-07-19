package email

import "fmt"

type Template struct {
	Body string
	Subject string
}

func WelcomeMail(username *string) *Template {
	return &Template{
		Body: fmt.Sprintf("Hey %s, welcome to our app, we are happy to have you onboard", *username),
		Subject: "Welcome to Receept!!!",
	}
}