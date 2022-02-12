package main

import (
	"context"
	"fmt"
	"runtime"
	"sort"
	"time"
)

func threeSum(nums []int) [][]int {
	//qsort(nums)
	//sort.Ints(nums)
	//n, ans := len(nums), make([][]int, 0)
	//
	//// 枚举 a
	//for first := 0; first < n; first++ {
	//    // 需要和上一次枚举的数不相同
	//    if first > 0 && nums[first] == nums[first - 1] {
	//        continue
	//    }
	//    // c 对应的指针初始指向数组的最右端
	//    third := n - 1
	//    target := -1 * nums[first]
	//    // 枚举 b
	//    for second := first + 1; second < n; second++ {
	//        // 需要和上一次枚举的数不相同
	//        if second > first + 1 && nums[second] == nums[second - 1] {
	//            continue
	//        }
	//        // 需要保证 b 的指针在 c 的指针的左侧
	//        for second < third && nums[second] + nums[third] > target {
	//            third--
	//        }
	//        // 如果指针重合，随着 b 后续的增加
	//        // 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
	//        if second == third {
	//            break
	//        }
	//        if nums[second] + nums[third] == target {
	//            ans = append(ans, []int{nums[first], nums[second], nums[third]})
	//        }
	//    }
	//}
	//return ans
	sort.Ints(nums)
	l, ans := len(nums), [][]int{}
	for a := 0; a < l-2; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		c := l - 1
		target := -nums[a]
		for b := a + 1; b < l-1; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return ans
}

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid, left, right := nums[0], 0, len(nums)-1
	if left >= right {
		return
	}
	for left < right {
		for left < right && nums[right] >= mid {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
		for left < right && nums[left] <= mid {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	qsort(nums[:left])
	qsort(nums[left+1:])
	ctx, cancel := context.WithCancel(context.Background()) //context.Background()返回空context
	go func(c context.Context) {
		i := 1
		for {
			fmt.Printf("打印%d次\n", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	runtime.Gosched()
}
