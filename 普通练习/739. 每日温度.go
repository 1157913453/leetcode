package main

//func dailyTemperatures(temperatures []int) []int {
//    res := make([]int, len(temperatures))
//    for i:=0;i < len(temperatures)-1;i ++{
//        for j := i+1;j<len(temperatures);j++{
//            if temperatures[j] > temperatures[i] {
//                res[i] = j-i
//                break
//            }
//        }
//    }
//    return res
//}

type Test struct {
	Name string
}

var t map[string]Test

func dailyTemperatures(tm []int) []int {
	t = make(map[string]Test)
	name := Test{"sjd"}
	t["Name"] = name
	x := t["Name"]
	x.Name = "sjk"
	l, prevIndex := len(tm), 0
	ans, stack := make([]int, l), []int{}
	for k, v := range tm {
		for len(stack) > 0 && v > tm[prevIndex] {
			prevIndex = stack[len(stack)-1]
			ans[prevIndex] = k - prevIndex
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
	}
	return ans
}
