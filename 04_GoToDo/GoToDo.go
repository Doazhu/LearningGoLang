package main

import "fmt"

func main() {
	var todolist = [...]string{
		"съесть авокадо",
		"найти квадрат",
		"почистить унитаз",
	}

	fmt.Println("по итогу задач: ", len(todolist))

	for index, item := range todolist {
		fmt.Printf("0%d, %s\n", index, item)
	}
}