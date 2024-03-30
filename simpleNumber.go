// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	l := 99
	r := make([]int, 0, 50)

	for i := 2; i < l; i++ {
		flag := 0
		for j := 2; j < l; j++ {
			if i%j == 0 {
				flag ++
				if flag == 2 {
					break
				}
			}

		}
		if flag == 1 {
			r = append(r, i)
			fmt.Println("Простое число: ", i)

		}
	}
	fmt.Println(r)

}
