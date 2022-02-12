package main

func maxSubArray(nums []int) int {
	sum, max := 0, nums[0]
	for _, v := range nums{
		if sum <= 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum{
			max = sum
		}
	}
	return max
}

func main() {

}
