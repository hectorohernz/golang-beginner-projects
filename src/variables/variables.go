package main

import "fmt"

func main() {

	// Static typed
	var a string = "string"
	fmt.Println(a)

	// Dynamic
	var x = "another string"
	fmt.Println(x)

	// Short hand version
	y := "Another Another String"
	fmt.Println(y)

	var b, c int = 1, 2
	fmt.Println(b, c)

	sum := 1 + 2
	fmt.Println(sum)

}
