package main

import "fmt"
// Для итерации по массивам, срезам и мапам используется следующая конструкция 
//for index, value := range array {} , где index будет указывать на текущий индекс
// элемента при каждой итерации, а value — значение элемента по этому же индексу.
func main(){
	var todolist = [3]string{}

	todolist[0] = "купить апельсины"
	todolist[1] = "купить гугл"
	todolist[2] = "купить яблоко"
	
	for index := range todolist{
		fmt.Printf("0%d. %s\n", index, todolist[index])
	}
}