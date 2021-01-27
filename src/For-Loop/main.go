package main

import "fmt"

func main() {
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for {
		fmt.Println("Loop")
		break
	}

	for j := 4; j <= 7; j++ {
		fmt.Println(j)
	}

	 x int32 = 1
	for x <= 3; x++ {
		fmt.Println(x)
	}
}
