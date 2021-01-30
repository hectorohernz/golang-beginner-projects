package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Println("Write", i, " as")

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	}

	t := time.Now().Weekday() // The Day of the week // Monday and .... Etc

	switch t {
	case time.Sunday, time.Saturday: // Using commas to seperate multiple expressions
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	whatInterfaceAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}

	whatInterfaceAmI(true)
	whatInterfaceAmI(1)
	whatInterfaceAmI("Hey")

}
