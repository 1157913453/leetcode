package main

import "fmt"

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	arr := []int{5, 3, 2}
	for _, i := range arr{
		for n%i == 0 {
			n /= i
		}
	}
	return n == 1
}

func main() {
	n := 6543124654
	fmt.Println(isUgly(n))
}
