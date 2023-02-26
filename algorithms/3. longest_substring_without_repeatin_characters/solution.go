package ___longest_substring_without_repeatin_characters

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	j := -1
	ans := 0
	m := make(map[byte]int)
	for i := 0; i < n; i++ {
		for j+1 < n && m[s[j+1]] == 0 {
			m[s[j+1]] = 1
			j++
		}
		newLen := j - i + 1
		if newLen > ans {
			ans = newLen
		}
		delete(m, s[i])
	}
	return ans
}
