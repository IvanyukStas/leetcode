package main

import (
	"fmt"
	"strings"
)

func main(){
	


	
	

}



func isPalindrome(s string) bool {
    s = strings.ToLower(s)
	var builder strings.Builder
	for i := 0; i < len(s); i++ {
		if isAlphanumeric(s[i]) {
			builder.WriteByte(s[i])
		}
	}
	s = builder.String()
	begin, end := 0, len(s)-1
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

func isAlphanumeric(c uint8) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}