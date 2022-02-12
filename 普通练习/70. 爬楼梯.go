package main

func climbStairs(n int) int {
	p, q, sum := 0, 0, 1
	for i:= 0;i < n;i++{
		p = q
		q = sum
		sum = q + p
	}
	return sum
}

func main() {

}
