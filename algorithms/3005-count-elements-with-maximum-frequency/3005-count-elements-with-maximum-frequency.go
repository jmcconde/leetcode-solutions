package leetcode

func maxFrequencyElements(nums []int) int {

	m := make(map[int]int)
	maxFrequency := 1
	result := 0

	for _, element := range nums {
		m[element] = m[element] + 1
		if m[element] > maxFrequency {
			maxFrequency = m[element]
		}
	}

	for _, element := range nums {
		if m[element] == maxFrequency {
			result += 1
		}
	}

	return result
}
