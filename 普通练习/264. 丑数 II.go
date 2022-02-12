package main

import "fmt"

//三指针法
//func nthUglyNumber(n int) int {
//	arr := []int{2,3,5}
//	q := 1
//	ans1, ans2, ans3 := 1, 1, 1
//	for i:=0;i<n;i++{
//
//
//	}
//}

func min(a, b, c int) int {
	if a<=b {
		if a<=c {
			return a
		} else {
			return c
		}
	} else {
		if b <=c {
			return b
		} else {
			return c
		}
	}
}



func main() {
	a, b, c := 7,5,0

	fmt.Println(min(a,b,c))
}
