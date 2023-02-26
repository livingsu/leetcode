package algorithms

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if idx, ok := m[target-num]; ok {
			return []int{i, idx}
		}
		m[num] = i
	}
	return []int{}
}
