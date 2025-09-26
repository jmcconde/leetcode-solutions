package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, element := range nums {
		val, isPresent := m[target-element]
		if isPresent {
			return []int{index, val}
		} else {
			m[element] = index
		}
	}
	return []int{nums[0], nums[1]}
}
