package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int
	fmt.Println(intNum)

	fmt.Println(len("AB"))
	fmt.Println(len("&J"))
	fmt.Println(len("@K"))

	fmt.Println(utf8.RuneCountInString("AB"))
	fmt.Println(utf8.RuneCountInString("&J"))
	fmt.Println(utf8.RuneCountInString("@K"))

	var myRune rune = 'a'
	fmt.Println(myRune)
}
