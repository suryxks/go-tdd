package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func Sum(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce[int](numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var ans []int
	for _, numbers := range numbersToSum {
		ans = append(ans, Sum(numbers))
	}

	return ans
}
func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}
	return Reduce[[]int](numbersToSum, sumTail, []int{})
}

func Reduce[T any, B any](collection []T, accumulator func(B, T) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64

	adjustBalance := func(currentBalnce float64, t Transaction) float64 {
		if t.From == name {
			return currentBalnce - t.Sum
		}
		if t.To == name {
			return currentBalnce + t.Sum
		}
		return currentBalnce
	}
	return Reduce(transactions, adjustBalance, balance)
}
