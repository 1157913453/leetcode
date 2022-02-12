package main

import "fmt"

//func getRow(rowIndex int) []int {
//	arr := make([][]int,rowIndex+1)
//	for i:=0 ;i<=rowIndex;i++ {
//		arr[i] = make([]int,i+1)
//		arr[i][0],arr[i][i] = 1,1
//		for n:= 1;n<i;n++{
//			arr[i][n] = arr[i-1][n-1] + arr[i-1][n]
//		}
//	}
//	return arr[rowIndex]
//}


//降低空间复杂度
func getRow(rowIndex int) []int {
	arr := make([]int,rowIndex+1)
	arr[0] = 1
	for i:= 1;i<=rowIndex;i++{
		for n:=i;n>0;n--{
			arr[n] += arr[n-1]
		}
	}
	return arr
}

func main() {
	fmt.Println(getRow(3))
}
