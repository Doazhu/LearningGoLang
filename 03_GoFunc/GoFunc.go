package main

import (
	"fmt"

	"errors"
)

const pi = 3.14

func main() {
	CalkAreaPrint(10)
	CalkAreaPrint(-9)
}

func CalkAreaPrint(radius int) {
	fmt.Println("Сап, формула pi*r^2, в переменой заначение: ", radius)
	calk, err := calkArea(radius)
	if err != nil{
		fmt.Println(err.Error())
		return 
	}

	fmt.Println("ответ: ", calk)
	fmt.Println("--------------------------------------")
}

func calkArea(radius int,) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("Переменая radius не может быть отрицательной") 
 	}

	return float32(radius) * float32(radius) * pi, nil
}
 
// 	var number1 int = 3
// 	var number2 int = 1

// func main() {

// 	fmt.Printf("Сумма выражения: ", plus)
// 	fmt.Println("Разность выражения: ", minus)
// }

// func plus(number1 int, number2 int) int {
// 	return int(number1) + int(number2)
// }

// func minus(number1 int, number2 int) uint {
// 	if number2 > number1{
// 		fmt.Println("change number 2")
// 		return
// 	}
// 	return number1 - number2 
// }

