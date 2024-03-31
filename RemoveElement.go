package main

import (
	"fmt"
)



func main(){

	nums := []int{3,2,2,3}
	val := 2
	res := removeElement(nums, val)
	fmt.Println(res,"--", nums)

	nums = []int{0,1,2,2,3,0,4,2}
	val = 2
	res = removeElement(nums, val)
	fmt.Println(res,"--", nums)
}

func removeElement(nums []int, val int) int {	
	k := len(nums)
	n := 0
	
    for i := 0; i <len(nums); i ++{
		if nums[i] == val{
			n++
		}
		if nums[i] == val && nums[k-1] != val {
			nums[i], nums[k-1] = nums[k-1], nums[i]
			k--
		}else if(nums[i] == val) && (nums[k-1] == val){
			for j := k-1; j > 0; j--{
				if nums[j] != val {
					nums[i], nums[j] = nums[j], nums[i]
					k--
					break
					
				}
			}
		}
	}
	return n
}