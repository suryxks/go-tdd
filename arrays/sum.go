package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var ans []int
	for _, numbers := range numbersToSum {
		ans = append(ans, Sum(numbers))
	}

	return ans
}
func SumAllTails(numbersToSum ...[]int) []int {
	var ans []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			ans = append(ans, 0)
		} else {
			tail := numbers[1:]
			ans = append(ans, Sum(tail))
		}

	}
	return ans
}
