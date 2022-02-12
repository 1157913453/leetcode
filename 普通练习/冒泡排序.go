package main

import "fmt"

func bubble_sor(nums []int){
	for i := len(nums)-1;i>=1;i--{
		for n := 0;n<i;n++{
			if nums[n] > nums[n+1]{
				nums[n], nums[n+1] = nums[n+1], nums[n]
			}
		}
	}
}

func main() {
	nums := []int{2,1}
	bubble_sor(nums)
	fmt.Println(nums)
}
