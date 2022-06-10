package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	for insertIndex := m + n - 1; insertIndex >= 0; insertIndex-- {
		if index1 < 0 {
			nums1[insertIndex] = nums2[index2]
			index2--
		} else if index2 < 0 {
			nums1[insertIndex] = nums1[index1]
			index1--
		} else if nums2[index2] > nums1[index1] {
			nums1[insertIndex] = nums2[index2]
			index2--
		} else {
			nums1[insertIndex] = nums1[index1]
			index1--
		}
	}
}
