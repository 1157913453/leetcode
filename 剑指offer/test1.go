//@Time : 2021/8/6 22:26
//@Author : zr
//@File : test1.go
//@Software : GoLand

package main

import "fmt"

func main() {
	defer func() {
		panic("90909")
	}()
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		fmt.Println("78787878")
	}()
	panic("111")
	panic("222")
}
