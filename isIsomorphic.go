package main

import (
	"fmt"
)

func main() {
	s := "egg"
	t := "add"
	fmt.Println(isIsomorphic(s, t))

	s = "\u0000\u0000"
	t = "\u0000c"
	fmt.Println(isIsomorphic(s, t))
}

	func isIsomorphic(s string, t string) bool {
		slcS := make([]int8, 255, 255)
		slcT := make([]int8, 255, 255)
		for i := range 255{
			slcS[i] = -1
			slcT[i] = -1
		}

		for i := 0; i < len(s); i++ {
			slcSIndx := int8(rune(s[i]))
			slcTIndx := int8(rune(t[i]))
			if slcS[slcSIndx] == -1 && slcT[slcTIndx] == -1 {
				slcS[slcSIndx] = slcTIndx
				slcT[slcTIndx] = slcSIndx
			} else if slcS[slcSIndx] != slcTIndx || slcT[slcTIndx] != slcSIndx {
				return false

			}

		}

		return true
	}
