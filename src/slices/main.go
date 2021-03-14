package main

import "fmt"

func main() {

	// Creating a slice of three elements
	s := make([]string, 3)

	// Setting the array
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("emp:", s)

	// Appending to the slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	//Creating a copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slicing support a slice
	l := s[1:3]
	fmt.Println("sl1:", l)

	l = s[:1]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
