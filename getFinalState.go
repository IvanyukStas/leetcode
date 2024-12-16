package main

import (
	"fmt"
)

func main(){

	arr := []int{2,1,3,5,6}
	fmt.Println(getFinalState(arr, 5, 2))


}

func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
		indMin := 0
		min := nums[0]
		flagEqual := 0

		for j := 1;  j < len(nums); j++ {
			if flagEqual == 1{
				continue
			}
			if min > nums[j] && flagEqual == 0{
				if min == nums[j]{
					flagEqual = 1
				}
				min = nums[j]
				indMin = j
			}
			
		}
		min = min*multiplier
			nums[indMin] = min
			
			flagEqual = 0
	}
	return nums
}