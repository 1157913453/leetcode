//@Time : 2021/7/30 13:02
//@Author : zr
//@File : 303. 区域和检索 - 数组不可变.go
//@Software : GoLand

package main
//  普通方法
//type NumArray struct {
//	nums []int
//}
//
//func Constructor(nums []int) NumArray {
//	return NumArray{nums: nums}
//}
//
//func (this *NumArray) SumRange(left int, right int) int {
//	if len(this.nums) == 0 {
//		return 0
//	}
//	sum := 0
//	for i:=left;i<=right;i++{
//		sum += this.nums[i]
//	}
//	return sum
//}


//  前缀和
//type NumArray struct {
//	sum []int
//}
//
//func Constructor(nums []int) NumArray {
//	sum :=make([]int,len(nums)+1)
//	for k, v := range nums{
//		sum[k+1] = sum[k] + v
//	}
//	return NumArray{sum}
//}

//func (this *NumArray) SumRange(left int, right int) int {
//	return this.sum[right+1] - this.sum[left]
//}



/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {

}
