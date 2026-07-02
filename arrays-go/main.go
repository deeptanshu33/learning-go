package main

func Sum(numbers []int) (res int) {
	for _, val := range numbers {
		res += val
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, slice := range numbersToSum {
		sums[i] = Sum(slice)
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, slice := range numbers {
		if len(slice) > 0 {
			sums[i] = Sum(slice[1:])
		}
	}

	return sums
}
