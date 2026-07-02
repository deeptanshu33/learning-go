package main

func Sum(numbers []int) (res int) {
	for _, val := range numbers {
		res += val
	}
	return
}
