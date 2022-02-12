//@Time : 2021/7/27 9:54
//@Author : zr
//@File : 莫比乌斯环.go
//@Software : GoLand

package main

import "fmt"

func main() {
	////a := 5  //用键盘输入队伍数量
	//q1 := 130  //用键盘输入
	//q2 := 210  // 用键盘输入
	//p := 100 //用键盘输入允许的最大分差
	//max := q1 + q2 + p
	//min := q1 + q2 - p
	////var b int = 5  //单人匹配玩家数量
	//var arr1 = [5][2]int{{100, 300}, {350, 30}, {190, 150}, {400, 100}, {230, 110}} //用键盘输入队伍数量
	//var arr2 = []int{200, 140, 200}                                                 //用键盘输入单人匹配的个人分数
	//sum := 0
	//for i := 0; i < len(arr1); i++ { //判断等待匹配的队伍中有多少符合要求
	//	if arr1[i][0]+arr1[i][1] >= min && arr1[i][0]+arr1[i][1] <= max {
	//		sum += 1
	//	}
	//}
	//for i := 0; i < len(arr2)-1; i++ {
	//	for n := i + 1; n < len(arr2); n++ {
	//		if arr2[i]+arr2[n] >= min && arr2[i]+arr2[n] <= max {
	//			sum += 1
	//		}
	//	}
	//}
	//fmt.Println(sum)
	//var N, T1, T2, T3, T4 int
	//var time = 0
	//var sw := 0  //切换次数
	//var arr [][]int
	//for i:=0;i<2;i++{
	//	for n:=0;n<N;n++{
	//		fmt.Scan("%d",arr[i][n])
	//	}
	//}

	var M, N, L int
	fmt.Scanln(&M, &N, &L)
	var arr = [2][]int{}      //初始化二维数组
	arr[0] = make([]int, N)
	arr[1] = make([]int, N)
	for i := 0; i < 2; i++ {       //从键盘向二维数组输入数据
		for n := 0; n < N; n++ {
			fmt.Scanf("%d", &arr[i][n])
			if i == 0 {                  //初始化所有小车距离起点的位置
				arr[0][n] = arr[0][n] % M
			}
		}
	}
	var tmp, time float32
	speed := arr[1][0]        //玩家小车速度
	distance := arr[0][0]    //玩家小车距离起点距离
	for i := 1; i < len(arr[0]); i++ {
		if arr[1][i] < speed {    //如果前车速度小于玩家
			tmp = float32(arr[0][i] - (distance + L)) / float32(speed - arr[1][i])
			if tmp > time{
				time = tmp
			}
		} else {   //如果前车速度大于玩家
			tmp = float32(M - arr[0][i] + distance - L) / float32(arr[1][i] - speed)
			if tmp > time{
				time = tmp
			}
		}
	}
	fmt.Printf("%.2f",time)
}
