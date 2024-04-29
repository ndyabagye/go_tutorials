package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = []rune("résumè")
	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\n The length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	// var catStr = ""
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catStr = strBuilder.String()

	fmt.Printf("\n%v", catStr)
}
