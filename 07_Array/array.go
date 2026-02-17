package main

import "fmt"

// func main() {
// 	var arr [4]int 

// 	fmt.Println(arr)
// } получим массив с длиной 4 символа

func main() {
	var arr [3]int

	fullArray(arr)

	fmt.Println(arr)
}

func fullArray(arr [3]int){
	for i := 0; i < len(arr); i++{
		arr[i] = i
	}

	fmt.Println("fullArray():", arr)
}
