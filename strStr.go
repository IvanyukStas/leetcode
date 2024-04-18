package main

import (
	"fmt"
)

func main(){
var 	haystack = "s1dbuts1ad"
	var needle = "sad"
	fmt.Println(strStr(haystack, needle))

	haystack = "mississippi"
	needle = "issi"
	fmt.Println(strStr(haystack, needle))

}


func strStr(haystack string, needle string) int {
    
	i := 0

	for {
		if i + len(needle) > len(haystack){
			break
		}
		if needle == haystack[i:i + len(needle)]{
			return i
		}else{
			i++
		}
	}

	return -1
}