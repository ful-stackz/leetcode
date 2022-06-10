package leetcode

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := -1

	for right > left {
		maxL := height[left]
		maxR := height[right]
		area := calcArea(height, left, right)
		if area > maxArea {
			maxArea = area
		}

		if maxL > maxR {
			for right > left {
				right--
				if height[right] > maxR {
					break
				}
			}
		} else {
			for right > left {
				left++
				if height[left] > maxL {
					break
				}
			}
		}
	}

	return maxArea
}

func calcArea(height []int, left int, right int) int {
	return min(height[left], height[right]) * abs(left-right)
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}
	return v
}
