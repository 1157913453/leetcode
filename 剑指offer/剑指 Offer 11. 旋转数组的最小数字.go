//@Time : 2021/7/26 10:12
//@Author : zr
//@File : 剑指 Offer 11. 旋转数组的最小数字.go
//@Software : GoLand

package main

import "fmt"

func minArray(numbers []int) int {
	for k, v := range numbers[1:] {
		if v < numbers[k] {
			return v
		}
	}
	return numbers[0]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	fmt.Println(minArray(nums))
}
