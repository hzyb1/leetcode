package main

import "leetcode/util"

func findNumberOfLIS(nums []int) int {
	dp := make([]int,len(nums))
	// maxNum := make([]int, len(nums))
	dp[0] = 1
	for i,_ := range nums {
		dp[i]=1
		for j:=0;j<i;j++ {
			if nums[j] < nums[i] {
				dp[i] = util.Max(dp[i], dp[j]+1)
			}
		}

	}
	ans := 0
	mx := dp[0]
	for _,num := range dp {
		mx = util.Max(mx, num)
	}
	ans = getNum(nums, dp, mx, len(nums)-1)
	if mx == 1 {
		return len(nums)
	}
	return ans
}

func getNum(nums, dp []int, mx ,pos int) int{
	if mx <=1 {
		return 1
	}
	ans := 0
	for i:=0;i<=pos;i++ {
		if mx == dp[i] {
			for j:=0;j<i;j++ {
				if nums[i]>nums[j] && dp[i] == dp[j]+1 {
					ans += getNum(nums, dp, mx-1, j)
				}
			}
		}
	}
	return ans
}