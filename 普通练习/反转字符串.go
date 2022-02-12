//@Time : 2021/7/8 22:48
//@Author : zr
//@File : 反转字符串.go
//@Software : GoLand

package main

func reverseString(s []byte)  {
	for i:=0;i<len(s)/2;i++{
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main()  {
}
//func rotate(nums []int, k int)  {
//	k=k%len(nums)
//	a:=nums[:len(nums)-k]
//	b:=nums[len(nums)-k:len(nums)]
//	nums = append(b,a...)
//
//}
//
//func main() {
//	nums:=[]int{1,2,3,4,5}
//	rotate(&nums,3)
//	fmt.Println(nums)
//}
