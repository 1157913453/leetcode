//@Time : 2021/7/30 14:40
//@Author : zr
//@File : 1365. 有多少小于当前数字的数字.go
//@Software : GoLand

package main

func smallerNumbersThanCurrent(nums []int) []int {
	arr:=make([]int,len(nums))
	for k, v := range nums{
		a:=0
		for i:=0;i<len(nums);i++{
			if v >nums[i] {
				a++
			}
		}
		arr[k] = a
	}
	
	return arr
}

func main() {

}
