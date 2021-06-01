package main

import "fmt"

func main() {
	x := make([]int, 10,100)
	fmt.Println( x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// 2D slice
	b1 := []string{"udith", "nalaka", "icecream"}
	b2 := []string{"iranga", "rana", "coffee"}

	fmt.Println(b1)
	fmt.Println(b2)

	b3 := [][]string{b1, b2}
	fmt.Println(b3)
}
