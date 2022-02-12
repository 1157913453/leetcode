package main

func quickSort(nums []int) {
	mid, left, right := nums[0], 0, len(nums)-1
	if left >= right {
		return
	}
	for left < right {
		for left < right && nums[right] >= mid {
			right--
		}

	}
}

func main() {

}
