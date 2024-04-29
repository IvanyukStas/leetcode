package main

import (
	"fmt"
	"strings"
)


func main(){
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println(wordPattern(pattern, s), "TRUE")

	pattern = "abba"
	s = "dog dog dog dog"
	fmt.Println(wordPattern(pattern, s), "FALSE")

	


}

func wordPattern(pattern string, s string) bool {
	hashmap :=  make(map[rune]string)
	slcS := strings.Split(s, " ")
	slcTmp := make([]string, len(slcS), len(slcS))
	if len(pattern) != len(slcS){
		return false
	}
	for i, v := range pattern{
		if value, ok := hashmap[v]; !ok {
			if len(slcTmp) > 0 {
				for _, val := range slcTmp{
					if val == slcS[i]{
						return false
					}
				}
			}
			hashmap[v] = slcS[i]
			slcTmp[i] = slcS[i]
		}else if slcS[i] != value{
			return false
		}		
	}
	return true
    
}