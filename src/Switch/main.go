package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Write", i, " as")

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}
}
