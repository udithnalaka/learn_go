package main

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersToSum)
	//sums := make([]int, lengthOfNumbers)

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTail(numbersToAdd ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToAdd {
		sums = append(sums, numbers[len(numbers)-1])
	}
	return sums
}

func removeFromSlice(number []int) []int {
	return append(number[:2], number[4:]...)
}

