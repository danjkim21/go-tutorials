package main

import "fmt"

func main() {

	var myArray [5]int 
	myArray[0] = 10
	myArray[4] = 15
	fmt.Println(myArray)

	mySlice := []int{3,5,6,7,2,1,32}
	mySlice = append(mySlice, 122)
	fmt.Println(mySlice)

	var myMap map[string]int = make(map[string]int)
	myMap["adam"] = 12
	myMap["bob"] = 13
	myMap["charlie"] = 14

	fmt.Println("adam is", myMap["adam"], "years old")
	fmt.Println(myMap)

	for i:=1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("FIVE")
			break
		}
		fmt.Println(i)
	}
}