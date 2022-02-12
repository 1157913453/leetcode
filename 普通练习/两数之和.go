package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, num := range nums {
		if v, exit := m[target - num]; exit {
			return []int{k, v}
		}
		m[num] = k
	}
	return nil
}
