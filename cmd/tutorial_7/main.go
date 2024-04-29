package main

import "fmt"

func main() {

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is :%p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
	// var slice = []int32{1, 2, 3}
	// var sliceCopy = slice

	// sliceCopy[2] = 4
	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	// var p *int32 = new(int32)
	// var i int32

	// fmt.Printf("The value p points to is: %v", *p)
	// fmt.Printf("\nThe value of i is: %v", i)
	// p = &i
	// *p = 1
	// fmt.Printf("\nThe value p points to now is: %v", *p)
	// fmt.Printf("\nThe value of i now is: %v", i)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing1 array is :%p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
