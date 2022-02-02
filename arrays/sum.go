package main

func Sum(a []int) int {
	sum := 0
	for _, nr := range a {
		sum += nr
	}
	return sum
}
