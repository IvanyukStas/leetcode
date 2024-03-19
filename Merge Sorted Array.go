package main

import (
	"fmt"
)

func main() {

	nums1 := []int{1, 2, 3, 0, 0}
	nums2 := []int{2, 4}
	m := 3
	n := 2
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums1 = []int{1}
	nums2 = []int{}
	m = 1
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	m = 0
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}

func merge(nums1 []int, m int, nums2 []int, n int) {
	t := 0
	if m < 1 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]

		}
	}
	if m > 0 && n > 0 {
		for i := m; i < m+n; i++ {
			nums1[i] = nums2[t]
			t++
		}
	}
	slices.Sort(nums1)
}
