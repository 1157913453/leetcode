//@Time : 2021/7/26 15:52
//@Author : zr
//@File : 整数反转.go
//@Software : GoLand

package main

import (
	"fmt"
)

//func reverse(x int) int {
//	arr := list.List{}
//	if x<0{
//		x = -1*x
//		for x > 9 {
//			arr.PushBack(x % 10)
//			x = x / 10
//		}
//		arr.PushBack(x)
//		sum := 0
//		for arr.Len() > 0 {
//			sum = sum*10 + arr.Remove(arr.Front()).(int)
//		}
//		if sum >2147483647{
//			return 0
//		}
//		return sum*-1
//	} else {
//		for x > 9 {
//			arr.PushBack(x % 10)
//			x = x / 10
//		}
//		arr.PushBack(x)
//		sum := 0
//		for arr.Len() > 0 {
//			sum = sum*10 + arr.Remove(arr.Front()).(int)
//		}
//		if sum >2147483647{
//			return 0
//		}
//		return sum
//	}
//}
//func reverse(x int) int {
//	len := make([]int, 0)
//	if x >= 0 {
//		for x > 9 {
//			len = append(len, x%10)
//			x = x / 10
//		}
//		len = append(len, x%10)
//		sum := 0
//		for _, v := range len {
//			sum = sum*10 + v
//		}
//		if sum > 2147483647 {
//			return 0
//		}
//		return sum
//	} else {
//		x = -1 * x
//		for x > 9 {
//			len = append(len, x%10)
//			x = x / 10
//		}
//		len = append(len, x%10)
//		sum := 0
//		for _, v := range len {
//			sum = sum*10 + v
//		}
//		if sum > 2147483647 {
//			return 0
//		}
//		return -1*sum
//	}
//
//}
//
func reverse(x int) int{
	sum := 0
	for x!=0{
		sum = sum *10 + x%10
		x /= 10
	}
	if sum>2147483648 || sum < -2147483648{
		return 0
	}
	return sum
}

func main() {
	fmt.Println(reverse(-1897600))
}
