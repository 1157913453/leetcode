//@Time : 2021/7/25 23:00
//@Author : zr
//@File : 剑指 Offer 10- I. 斐波那契数列.go
//@Software : GoLand

package main

import "fmt"

func fib(n int) int {
	a, b := 0, 1
	for i:=0;i<n;i++{
		a, b = b, (a+b)%1000000007
	}
	return a
}

func main() {
	fmt.Println(fib(66))

}
