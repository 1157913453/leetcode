//@Time : 2021/7/27 0:56
//@Author : zr
//@File : 验证回文串.go
//@Software : GoLand

package main

import (
	"fmt"
	"strings"
)

//func isPalindrome(s string) bool { // 内存占用太多
//	arr := make([]int, len(s))
//	n := 0
//	for _, v := range s {
//		if v >= 65 && v <=90 && v > 47 && v < 58{
//			arr[n] = int(v)
//			n++
//		} else if v >= 97 && v <= 122 {
//			arr[n] = int(v-32)
//			n++
//		}
//	}
//	fmt.Println(arr)
//	for i:=0;i<n/2;i++{
//		if arr[i] != arr[n-i-1] {
//			return false
//		}
//	}
//	return true
//}

//func isPalindrome(s string) bool { //  用双指针做
//	s = strings.ToLower(s)
//	head, tail := 0, len(s)-1
//	for head < tail {
//		for head < tail && !(s[head] > 96 && s[head] < 123 || (s[head] > 47 && s[head] < 58)) {
//			head++
//		}
//		for head < tail && !(s[tail] > 96 && s[tail] < 123 || (s[tail] > 47 && s[tail] < 58)) {
//			tail--
//		}
//		fmt.Println(s[head],s[tail])
//		if head < tail {
//			if s[head] != s[tail]{
//				return false
//			}
//			head++
//			tail--
//		}
//	}
//	return true
//}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
