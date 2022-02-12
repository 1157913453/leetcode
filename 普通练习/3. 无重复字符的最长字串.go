package main

func lengthOfLongestSubstring(s string) int {
	m := map[int32]bool{}
	left, res, tmp := 0, 0, 0
	for k, v := range s {
		if m[v] == true {
			res = max(res, tmp)
			for ; s[left] != uint8(v); left++ {
				m[int32(s[left])] = false
			}
			left++
		} else {
			m[v] = true
			tmp = k - left + 1
		}
	}
	res = max(res, tmp)
	return res
}

//func lengthOfLongestSubstring(s string) int {
//	p, q, result, l := 0, 0, 0, len(s)
//	m := make(map[byte]int)
//	for ; q < l; q++ {
//		m[s[q]]++
//		for m[s[q]] > 1 {
//			m[s[p]]--
//			p++
//		}
//		result = max(result, q-p+1)
//	}
//	return result
//}
//
//func lengthOfLongestSubstring(s string) int {
//	m := map[byte]int{}
//	rk, ans, n := -1, 0, len(s)
//	for i := 0; i < n; i++ {
//		if i != 0 {
//			delete(m, s[i-1])
//		}
//		for rk+1 < n && m[s[rk+1]] == 0 {
//			m[s[rk+1]]++
//			rk++
//		}
//		ans = max(ans, rk-i+1)
//	}
//	return ans
//}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
