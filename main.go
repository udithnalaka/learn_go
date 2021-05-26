package main

import (
	"fmt"
	"math"

	"github.com/udithnalaka/learn_go/01_hello/strutil"
)

func main() {

	////data types
	//var name string = "Nalaka"
	name := "Nalaka"
	email := "u@u.com.au"
	var age int = 40
	const isMale = true

	fmt.Println("Hello Udith " + name + "!")

	fmt.Println(age)
	fmt.Println(email)
	fmt.Printf("%T\n", isMale)

	////math package
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	////util functions
	fmt.Println(strutil.Reverse("Udith"))
	fmt.Println(strutil.Greeting("Udith Nalaka"))

	fmt.Println(strutil.Add(3, 7))
	fmt.Println(strutil.Multiply(3, 7))

	////Arrays
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"
	fruitArr := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grapes"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1])

	////for loop
	i := 1
	for i <= 10 {
		fmt.Print(i)
		i++
	}
	fmt.Print("\n")

	////fizzbuzz
	FizzBuzz()
}

func FizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
