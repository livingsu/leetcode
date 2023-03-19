package _3__Maximum_Subarray

// 1. dynamic programming
func maxSubArray_dp(nums []int) int {
	n := len(nums)
	ans := nums[0]
	prevSum := nums[0]
	for i := 1; i < n; i++ {
		prevSum = max(prevSum+nums[i], nums[i])
		ans = max(ans, prevSum)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 2. divide and conquer
func maxSubArray_dac(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	mid := l + (r-l)/2
	lStatus := get(nums, l, mid)
	rStatus := get(nums, mid+1, r)
	return pushUp(lStatus, rStatus)
}

func pushUp(l, r Status) Status {
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	iSum := l.iSum + r.iSum
	mSum := max(l.mSum, max(r.mSum, l.rSum+r.lSum))
	return Status{lSum, rSum, iSum, mSum}
}

type Status struct {
	lSum, rSum, iSum, mSum int
}
