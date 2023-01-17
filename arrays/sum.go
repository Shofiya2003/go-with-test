package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var res []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			tail := numbers[1:]
			res = append(res, Sum(tail))
		}

	}

	return res
}
