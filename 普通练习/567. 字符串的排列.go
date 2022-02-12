package main

//func checkInclusion(s1 string, s2 string) bool {
//	l1, l2 := len(s1), len(s2)
//	if l1 > l2 {
//		return false
//	}
//	a := [26]int{}
//	for _, v := range s1 {
//		a[v-'a']--
//	}
//	left := 0
//	for right, v := range s2 {
//		x := v - 'a'
//		a[x]++
//		for a[x] > 0 {
//			a[s2[left]-'a']--
//			left++
//		}
//		if right - left == l1 {
//			return true
//		}
//	}
//	return false
//}

func checkInclusion(s1 string, s2 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1 > l2 {
		return false
	}
	diff := 0
	cnt := [26]int32{}
	for k, v := range s1 {
		cnt[v-'a']--
		cnt[s2[k]-'a']++
	}
	for _, v := range cnt {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := l1; i < l2; i++ {
		x, y := s2[i] - 'a', s2[i-l1] - 'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
            diff++
        }
        cnt[y]--
        if cnt[y] == 0 {
            diff--
        }
		if diff == 0 {
			return true
		}
	}


	return false
}
