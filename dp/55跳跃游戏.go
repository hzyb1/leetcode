package dp
func CanJump(nums []int) bool {
	if len(nums) <=1 {
		return true
	}
	n := len(nums)
	po,lo := 0,0
	for i:=0;i<n-1;{
		j:=i
		for ;j<n;j++ {
			if nums[j] == 0 {
				break
			}
		}
		if j==n {
			return true
		}
		po = j
		lo = 0
		for ;j<n;j++ {
			if nums[j] != 0 {
				break
			}
			lo++
		}
		k:=po-1
		for ;k>=0;k-- {
			if k+nums[k]+1 > lo+po || (k+nums[k]+1 >= lo+po && lo+po == n){
				break
			}
		}
		if k<0 {
			return false
		}
		i = po+lo
	}
	return true
}
