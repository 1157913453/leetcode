//@Time : 2021/7/30 15:27
//@Author : zr
//@File : 304. 二维区域和检索 - 矩阵不可变.go
//@Software : GoLand

package main

//  暴力法，超出时间限制
//type NumMatrix struct {
//	num [][]int
//}
//
//
//func Constructor(matrix [][]int) NumMatrix {
//	return NumMatrix{matrix}
//}
//
//
//func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
//	sum :=0
//	for i:=row1;i<=row2;i++{
//		for n :=col1;n<=col2;n++{
//			sum += this.num[i][n]
//		}
//	}
//	return sum
//}

//  一维前缀和
//type NumMatrix struct {
//	nums [][]int
//}
//
//func Constructor(matrix [][]int) NumMatrix {
//	arr := make([][]int,200)
//	//初始化二维数组
//	for i := range arr{
//		arr[i] = make([]int,201)
//	}
//	for k, v := range matrix{
//		for i:=0;i<len(v);i++{
//			arr[k][i+1] = arr[k][i] + v[i]
//		}
//	}
//	return NumMatrix{arr}
//}
//
//func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
//	sum1:=0
//	for i:=row1;i<=row2;i++{
//		sum1 += this.nums[i][col2+1]- this.nums[i][col1]
//	}
//	return sum1
//}

//  二维前缀和

type NumMatrix struct {
	nums	[][]int
}


//func Constructor(matrix [][]int) NumMatrix {
//	arr := make([][]int,200)
//
//
//}


//func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
//
//}

func main() {

}
