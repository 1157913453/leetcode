//@Time : 2021/7/29 21:35
//@Author : zr
//@File : 剑指 Offer 42. 连续子数组的最大和.go
//@Software : GoLand

package main

import (
	"fmt"
)

//func maxSubArray(nums []int) int {
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	sum := nums[0]
//	max := sum
//	if sum >= 0 {
//		for i := 1; i < len(nums); i++ {
//			if sum+nums[i] > 0 {
//				sum += nums[i]
//				max = MaxNum(sum,max)
//			} else {
//				sum = 0
//				max = MaxNum(sum,max)
//			}
//		}
//	} else {
//		for i := 1; i < len(nums); i++ {
//			if sum > 0 {
//				if sum + nums[i]>0{
//					sum += nums[i]
//					max = MaxNum(sum,max)
//				} else {
//					sum = 0
//					max = MaxNum(sum,max)
//				}
//			} else {
//				if nums[i] > 0 {
//					sum = nums[i]
//					max = MaxNum(sum,max)
//				} else if nums[i] > sum {
//					sum =nums[i]
//					max = MaxNum(sum,max)
//				}
//			}
//		}
//	}
//	return max
//}
//
//func MaxNum(a, b int) int {
//	if a>b{
//		return a
//	} else {
//		return b
//	}
//}


func maxSubArray(nums []int) int {
	min, max, sum := 0, nums[0],0
	for _, v := range nums{
		sum += v
		if sum - min > max{
			max = sum - min
		}
		if sum < min {
			min = sum
		}
	}
	return max
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}
