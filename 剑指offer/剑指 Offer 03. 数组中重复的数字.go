//@Time : 2021/7/27 18:41
//@Author : zr
//@File : 剑指 Offer 03. 数组中重复的数字.go
//@Software : GoLand

package main

func findRepeatNumber(nums []int) int {
	arr := make([]int, len(nums))
	for _, v := range nums {
		arr[v]++
	}
	for k, v := range arr {
		if v >=2 {
			return k
		}
	}
	return -1
}

func main() {

}
