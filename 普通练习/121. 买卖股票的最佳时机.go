package main

func maxProfit(prices []int) int {
	min, sum := prices[0], 0
	for _, v:= range prices {
		if min > v {
			min = v
		} else if v-min > sum{
			sum = v-min
		}
	}
	return sum
}

func main() {
	//prices = []int{7,1,5,3,6,4}

}
