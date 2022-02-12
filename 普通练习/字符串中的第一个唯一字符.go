//@Time : 2021/7/26 17:28
//@Author : zr
//@File : 字符串中的第一个唯一字符.go
//@Software : GoLand

package main

//func firstUniqChar(s string) int { //  哈希表方法
//	mp := make(map[int]int)
//	for _, v := range s {
//		mp[int(v)]++
//	}
//	for k, v := range s{
//		if mp[int(v)] == 1{
//			return k
//		}
//	}
//	return -1
//}

//func firstUniqChar(s string) int { //  数组方法更快
//	arr := make([]int, 26) //  26个字母
//	for _, v := range s {
//		arr[v-97]++
//	}
//	for k, v := range s {
//		if arr[v-97] == 1{
//			return k
//		}
//	}
//	return -1
//}

func firstUniqChar(s string) int {
	arr := make([]int, 27)
	for k, v := range s {
		if arr[v- 97] == 0 {  // 如果该字母之前没出现过，则值等于索引+1，因为索引从0开始，为了区分a[0]
			arr[v- 97] = k+1
		} else if arr[v-97] > 0 {
			arr[v-97] = -1
		}
	}

	//  查找arr中值的最小值
	res := -1
	for _, v:= range arr{
		if v>0{  //  说明有出现了一次
			if res <0{  //
				res = v-1
			} else if res+1 > v{
				res = v-1
			}
		}
	}
	return res
}

func main() {
	str := "asdasdasdasd"
	firstUniqChar(str)
}
