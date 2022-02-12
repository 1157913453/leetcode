//@Time : 2021/7/18 14:43
//@Author : zr
//@File : 旋转数组.go
//@Software : GoLand

package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse1(nums)
	reverse1(nums[:k])
	reverse1(nums[k:])
}

func reverse1(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

}
