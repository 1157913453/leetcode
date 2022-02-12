package main

import (
	"fmt"
	"math"
)

func main() {
	//chan1 := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(2)
	//var number int
	//for i:=0;i<2;i++{
	//	go func() {
	//		for {
	//			temp := <- chan1
	//			number = temp
	//			if number > 100 {
	//				fmt.Println("退出")
	//				wg.Done()
	//				return
	//			} else {
	//				number += 1
	//				temp = number
	//				chan1 <- temp
	//			}
	//		}
	//	}()
	//}
	//chan1 <- 0
	//wg.Wait()
	var n int
	var temp1 = 1
	var temp2 int
	var sum = 2
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println(n)
	}
	for i := 2; i < n; i++ {
		temp2 = temp1
		temp1 = sum
		sum = (temp2 + sum) % (int(math.Pow10(9)) + 7)
	}
	fmt.Println(sum)
	//n := 0
	//fmt.Scan(&n)
	//arr := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&arr[i])
	//}
	//m := make(map[int]bool)
	//for _, v := range arr {
	//	if m[v] == true {
	//		fmt.Println("NO")
	//		return
	//	}
	//	m[v] = true
	//}
	//fmt.Println("YES")
}
