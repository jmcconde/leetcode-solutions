package leetcode

// O(n) in-place solution
func MoveZeroes(nums []int) {
	pointer := 0
	for index, element := range nums[1:] {
		if nums[pointer] != 0 {
			pointer++
		} else if element == 0 {
			continue
		} else {
			nums[pointer] = element
			nums[index+1] = 0
			pointer++
		}
	}
}
