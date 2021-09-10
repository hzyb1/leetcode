package dfs

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := [][]int{}
	tmp := []int{}
	n := len(candidates)
	var dfs func(pos, target int)
	dfs = func (pos, target int) {
		target -= candidates[pos]
		tmp = append(tmp, candidates[pos])
		defer func() {tmp = tmp[:len(tmp)-1] }()
		if target == 0 {
			cp := make([]int, len(tmp))
			copy(cp, tmp)
			ans = append(ans, cp)
		} else if target < 0 {
			return
		}
		for i:=pos+1;i<n;i++ {
			dfs(i, target)
			j:=i+1
			for ;j<n;j++ {
				if candidates[pos]!=candidates[j] {
					break
				}
			}
			i = j
		}
	}
	for i:=0;i<n; {
		dfs(i, target)
		j:=i+1
		for ;j<n;j++ {
			if candidates[i]!=candidates[j] {
				break
			}
		}
		i = j
	}
	return ans
}
