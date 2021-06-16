package findMedianSortedArrays

import (
	"math"
	"sort"
)

//Complexity does not reach O(log (m+n))
func fakeFindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	tmp := []int{}
	tmp = append(nums1, nums2...)
	sort.Ints(tmp)
	l := len(tmp)
	if l%2 == 1 {
		return float64(tmp[int(l/2)])
	} else {
		t := int(len(tmp) / 2)
		return (float64(tmp[t-1]) + (float64(tmp[t]))) / 2
	}
}

//real solution
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//Make nums1 a shorter array, which can improve the retrieval speed and also avoid some boundary problems
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	l1, l2 := len(nums1), len(nums2)
	k := (l1 + l2 + 1) / 2 //After merging and sorting the two arrays, the length of the left half
	//Dichotomous search on nums1
	left := 0
	right := l1
	var result float64
	for left <= right {
		/*
		 * The position of the two arrays of measured numbers A,B (counting from 1)
		 * if m1 = 2 means the second number of the num1 array is 1 greater than index
		 */
		m1 := left + (right-left)/2
		m2 := k - m1
		if m1 > 0 && nums1[m1-1] > nums2[m2] {
			//A is bigger than B.next
			right = m1 - 1
		} else if m1 < l1 && nums1[m1] < nums2[m2-1] {
			//B is bigger than A.next
			left = m1 + 1
		} else {
			//Get the median
			if m1 == 0 {
				result = float64(nums2[m2-1]) //When the number of num1 array are on the right side of the total group
			} else if m2 == 0 {
				result = float64(nums1[m1-1]) //When the number of num2 array are on the left side of the total group
			} else {
				result = math.Max(float64(nums1[m1-1]), float64(nums2[m2-1]))
			}
			if (l1+l2)%2 == 1 {
				return result
			} else {
				var nextValue float64
				if m1 == l1 {
					nextValue = float64(nums2[m2])
				} else if m2 == l2 {
					nextValue = float64(nums1[m1])
				} else {
					nextValue = math.Min(float64(nums1[m1]), float64(nums2[m2]))
				}
				return (result + nextValue) / 2
			}
		}
	}
	return 0.0
}
