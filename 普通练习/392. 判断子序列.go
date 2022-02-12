package main

//func isSubsequence(s string, t string) bool {
//	t_len, s_len, k :=  len(t), len(s), 0
//	if s_len == 0{
//		return true
//	}
//	for i := 0; i < t_len; i++ {
//		if t[i] == s[k] {
//			if k == s_len - 1{
//				return true
//			}
//			k++
//		}
//	}
//	return false
//}

//双指针法
func isSubsequence(s string, t string) bool {
	sLen , tLen := len(s), len(t)
	j:= 0
	for i:=0;i<tLen;i++{
		if t[i] == s[j]{
			j++
		}
		if j == sLen{
			return true
		}
	}
	return false
}

func main() {

}
