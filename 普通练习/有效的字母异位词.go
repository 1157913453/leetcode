//@Time : 2021/7/26 20:02
//@Author : zr
//@File : 有效的字母异位词.go
//@Software : GoLand

package main

import "fmt"

//func isAnagram(s string, t string) bool {   //引用sort包进行排序
//	if len(s) != len(t) {
//		return false
//	}
//	sArr, tArr := make([]int,len(s)), make([]int,len(s))
//	for k, v := range s {
//		sArr[k] = int(v)
//	}
//	for k, v := range t {
//		tArr[k] =  int(v)
//	}
//	sort.Sort(sort.IntSlice(tArr))
//	sort.Sort(sort.IntSlice(sArr))
//	for i :=0;i<len(tArr);i++{
//		if sArr[i] != tArr[i]{
//			return false
//		}
//	}
//	return true
//}
func isAnagram(s string, t string) bool{
	if len(s) != len(t){
		return false
	}
	a := [26]int{}
	for _, v:= range s {
		a[v-97]++
	}
	for _, v:= range t {
		a[v-97]--
	}
	for i:=0;i<26;i++ {
		if a[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s,t))
}
