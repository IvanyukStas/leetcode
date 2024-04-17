package main

import (
	"fmt"
)

func main(){
	s := "a"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
	if len(s) == 1{
		return 1
	}
	 count:= 0
	 i := len(s)-1
	for {		
		if string(s[i]) != " "{			
			count++
		}else if string(s[i]) == " " && count > 0{			
			return count
		}
		if i == 0{
			return count
		}
		i--
	}
}