package greetings

import "fmt"

func Hello(Name string) string {
	massage := fmt.Sprintf("Привет, %v. Добро пожаловать!", Name)
	return massage
}
