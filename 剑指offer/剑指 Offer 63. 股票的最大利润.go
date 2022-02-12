//@Time : 2021/7/27 17:14
//@Author : zr
//@File : 剑指 Offer 63. 股票的最大利润.go
//@Software : GoLand

package main

import "fmt"

//func maxProfit(prices []int) int {  // 双指针穷举法
//	sum :=0
//	for i:=0;i<len(prices)-1;i++{
//		for n:= i+1;n<len(prices);n++{
//			if c := prices[n] - prices[i];c >sum{
//				sum = c
//			}
//		}
//	}
//	return sum
//}

func maxProfit(prices []int) int{
	if len(prices) <= 1{
		return 0
	}
	minPrice, money := prices[0], 0
	for i:=1;i<len(prices);i++ {
		if prices[i] - minPrice > money {
			money = prices[i] - minPrice
		}
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return money
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}
