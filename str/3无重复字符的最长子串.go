package str

import "leetcode/util"

func LengthOfLongestSubstring(s string) int {
	mp := map[byte]int{}
	l,r := 0, 0
	ans := 0
	for l<=r && r<len(s) {
		if mp[s[r]] >= l {
			ans = util.Max(ans, r-l)
			l = mp[s[r]]
		}
		mp[s[r]] = r+1
		r++
	}
	ans = util.Max(ans, r-l)
	return ans
}
