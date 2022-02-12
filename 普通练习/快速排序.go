//@Time : 2021/7/18 15:52
//@Author : zr
//@File : 快速排序.go
//@Software : GoLand

package main

import (
	"fmt"
)

func QSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	first, mid, last := 0, arr[0], len(arr)-1
	//rand.Seed(time.Now().Unix())
	//flag := rand.Intn(last)
	//arr[flag], arr[first] = arr[first], arr[flag]
	for first < last {
		for first < last && arr[last] >= mid {
			last--
		}
		arr[first], arr[last] = arr[last], arr[first]
		for first < last && arr[first] <= mid {
			first++
		}
		arr[first], arr[last] = arr[last], arr[first]
	}
	QSort(arr[:first])
	QSort(arr[first+1:])
}

//单路
//func QSort(arr []int, first, last int) {
//	if first >= last {
//		return
//	}
//	rand.Seed(time.Now().Unix())
//	k := rand.Intn(last-first) +first
//	arr[k], arr[first] = arr[first], arr[k]
//	k = first
//	for i:=first;i<last;i++{
//		if arr[last] > arr[i]{
//			arr[k], arr[i] = arr[i], arr[k]
//			k++
//		}
//	}
//	arr[k], arr[last] = arr[last], arr[k]
//	QSort(arr,first,k-1)
//	QSort(arr,k+1,last)
//}

//func QSort(arr []int, first, last , k int) int {
//	mid, left, right := first, first, last
//	ans := 0
//	if left >= right {
//		return arr[ans]
//	}
//	for first < last {
//		for first < last && arr[last] >= arr[mid] {
//			last--
//		}
//		arr[first], arr[last] = arr[last], arr[first]
//		mid = last
//		for first < last && arr[first] <= arr[mid]{
//			first++
//		}
//		arr[first], arr[last] = arr[last], arr[first]
//		mid = first
//	}
//	if len(arr)-mid == k{
//		ans = mid
//		return arr[ans]
//	} else if len(arr)-mid < k {
//		QSort(arr,left, mid-1,k)
//	} else {
//		QSort(arr,mid+1,right,k)
//	}
//	return arr[ans]
//	//QSort(arr,left, mid-1)
//	//QSort(arr,mid+1, right)
//}
func main() {
	b := []int{-1,9,9,5,8,5,4,3,6,7,8,6,7,0,1,2,-1,-4}
	QSort(b)
	fmt.Println(b)
}
