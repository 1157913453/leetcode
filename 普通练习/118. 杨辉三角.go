package main

func generate(numRows int) [][]int {
	arr := make([][]int,numRows)
	for i:=0;i<numRows;i++{
		arr[i] = make([]int,i+1)
		arr[i][0], arr[i][i] = 1, 1
		for n := 1;n<i;n++{
			arr[i][n] = arr[i-1][n-1] + arr[i-1][n]
		}
	}
	return arr
}

func main() {

}
