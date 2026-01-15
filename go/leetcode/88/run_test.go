package _test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var m = 3
	var nums2 = []int{2, 5, 6}
	var n = 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	var result []int
	var i, j int
	for i < m || j < n {
		if j >= n || (i < m && nums1[i] <= nums2[j]) {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	var count = len(nums1)
	for k := 0; k < count; k++ {
		nums1[k] = result[k]
	}
}
