//@Time : 2021/7/13 23:56
//@Author : zr
//@File : 旋转图像.go
//@Software : GoLand

package main

import "fmt"

func rotate1(matrix [][]int) {
	len := len(matrix[0])-1
	for i := 0; i < (len+1)/2; i++ {
		for n := 0; n < len-i*2; n++ {
			temp := matrix[i][n]
			matrix[i][n] = matrix[len-n][i]
			matrix[len-n][i] = matrix[len-i][len-n]
			matrix[len-i][len-n] = matrix[n][len-i]
			matrix[n][len-i] = temp
		}
	}
}

func main() {
	matrix := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate1(matrix)
	fmt.Println(matrix)
}
