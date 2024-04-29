package main

import (
	"fmt"
)

func main(){
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println(canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	hashmap := make( map[rune]int)

	for _, v := range magazine {
		if _, ok := hashmap[v]; !ok{
			hashmap[v] = 1
		}else{
			hashmap[v]++
		}
	}
	for _, v := range ransomNote {
		val, ok := hashmap[v]
		if ok && val > 0 {
			hashmap[v]--
		}else{
			return false
		}
	}
    return true
}