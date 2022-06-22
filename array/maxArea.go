package array

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	heightLen := len(height)
	var start, end = 0, heightLen - 1
	maxValue := 0
	for start < end {
		startValue := height[start]
		endValue := height[end]
		tempMax := (end - start) * min(startValue, endValue)
		if tempMax > maxValue {
			maxValue = tempMax
		}
		if startValue > endValue {
			end--
		} else {
			start++
		}
	}
	return maxValue
}
