package main

import (
	"fmt"
)

func main() {
	// nums := []int{1,1,1,1,1,2,2,3,3,4}
	nums := []int{1,}

	res := removeDuplicates(nums)
	fmt.Println(res)
	fmt.Println(nums)

}

func removeDuplicates(nums []int) int {
	switch{
	case len(nums) == 0:
		return 0
	case len(nums) == 1:
		return 1
	}
	}
	count := 0
	newMums := make([]int, len(nums))
	copy(newMums, nums)
	for i, v := range nums {
		if i == len(nums) - count{
			break
		}
		j := i+1
		for {
			if nums[j] == v{
				count++
				for x := j; x < len(nums)-count; x++{
					nums[x], nums[x+1]= nums[x+1], nums[x]
				}		
				
			} else{
				j++
			}
			if j >= len(nums) - count{
				break
			}
		}
		
	}
	return len(nums[:len(nums)-count])
}
