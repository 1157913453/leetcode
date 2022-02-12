//@Time : 2021/7/25 23:12
//@Author : zr
//@File : 剑指 Offer 10- II. 青蛙跳台阶问题.go
//@Software : GoLand

package main

import "fmt"

func numWays(n int) int {
	a, b := 1, 1
	for i:=0;i<n;i++{
		a, b = b, (a+b)%1000000007
	}
	return a
}
func main() {
	fmt.Println(numWays(3))
}
