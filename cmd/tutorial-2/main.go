package main

import "fmt"

func main() {

	var myArray [5]int 
	myArray[0] = 10
	myArray[4] = 15

	for i, value := range myArray {
		fmt.Println(i, value)
	}

	mySlice := []int{3,5,6,7,2,1,32}
	mySlice = append(mySlice, 122)

	for i, value := range mySlice {
		fmt.Println(i, value)
	}

	var myMap map[string]int = make(map[string]int)
	myMap["adam"] = 12
	myMap["bob"] = 13
	myMap["charlie"] = 14

	fmt.Println("adam is", myMap["adam"], "years old")

	for i, value:= range myMap {
		fmt.Printf("%v is %v years old\n", i, value)
	}

	for i:=1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("FIVE")
			break
		}
		fmt.Println(i)
	}
}