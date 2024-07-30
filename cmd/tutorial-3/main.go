package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "Hello World!"
	var indexedStr = myString[0:5]
	var indexedLetter = myString[0]

	fmt.Println(myString, indexedStr , indexedLetter)

	// for i, value := range myString {
	// 	fmt.Println(i, value)
	// }

	var strSlice = []string{"H", "e", "l", "l", "o", " ", "W", "o", "r", "l", "d", "!"}
	var catSlice = ""

	for i := range strSlice {
		catSlice += strSlice[i]
	}

	fmt.Println(catSlice)

	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catString = strBuilder.String()

	fmt.Println(catString)
}