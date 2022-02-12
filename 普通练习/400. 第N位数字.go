package main

import (
	"fmt"
	"math"
)

func findNthDigit(n int) int {
	d := 1
	for count := 9; n > d*count; count *= 10 {
		n -= d * count
		d++
	}

	start := math.Pow10(d - 1)
	num := int(start) + (n-1)/d
	dingnum := (n - 1) % d
	return num / int(math.Pow10(d-dingnum-1)) % 10
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Println(findNthDigit(n))
}
