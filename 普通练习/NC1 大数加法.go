package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	// write code here
	ls, lt, x, y, carry, sum, char := len(s)-1, len(t)-1, 0, 0, 0, 0, ""
	for ls >= 0 || lt >=0 || carry != 0 {
		if ls < 0 {
			x = 0
		} else {
			x = int(s[ls] - '0')
		}
		if lt < 0 {
			y = 0
		} else {
			y = int(t[lt] - '0')
		}
		ls--
		lt--
		sum = x + y + carry
		char = string(sum%10+48) + char
		carry = sum/10
	}
	return char
}
func main() {
	fmt.Println(solve("4346","231"))
}
