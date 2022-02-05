package main

func Sum(a []int) int {
	sum := 0
	for _, nr := range a {
		sum += nr
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, nrs := range numbersToSum {
		sums = append(sums, Sum(nrs))
	}
	return sums
}

func SumTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nrs := range numbersToSum {
		if len(nrs) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nrs[1:]))
		}
	}
	return sums
}
