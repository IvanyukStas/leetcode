// find max profit buy min sell max
package main

import (
	"fmt"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	r := maxProfit(prices)
	fmt.Println(r)

	prices1 := []int{3, 2, 1}
	r1 := maxProfit(prices1)
	fmt.Println(r1)

	prices2 := []int{2, 1, 4}
	r2 := maxProfit(prices2)
	fmt.Println(r2)

	prices21 := []int{2, 4, 1}
	r21 := maxProfit(prices21)
	fmt.Println(r21)

	prices211 := []int{2,1,2,1,0,1,2}
	r211 := maxProfit(prices211)
	fmt.Println(r211)

	prices2111 := []int{3,3,5,0,0,3,1,4}
	r2111 := maxProfit(prices2111)
	fmt.Println(r2111)

	prices21111 := []int{3,2,6,5,0,3}
	r21111 := maxProfit(prices21111)
	fmt.Println(r21111)

}


func maxProfit(prices []int) int {
	min := prices[0]
	max := 0
	var  res int
	for _, v := range prices[1:]{
		if min > v {
			// tmpMin = min
			min = v
			// tmpMax = max
			max = 0
		}else if max < v{
			max = v
		}
		if max - min > res{
			res = max- min
		}
	}
	fmt.Printf("min: %v " + " max: %v            ",min,  max)
	return res
	// if max == 0 && tmpMax >0 && tmpMin > 0 {
	// 	return tmpMax - tmpMin
	// }
	// res := max - min
	// if res < 1{
	// 	return 0
	// } else{
	// 	return res
	// }
}


func maxProfit1(prices []int) int {
	p := &ProductPrice{
		min:     prices[0],
		indxmin: 0,
	}

	for i, v := range prices {
        
		if p.min > v {			
			p.tempMin = p.min
			p.min = v
			p.indxmin = i
		} else if v >= p.max && i > p.indxmin{
			p.max = v
			p.indxmax = i
		}

	}
	if p.max == 0{
		p.tempMin = 0
	}
	fmt.Printf("%#v", p)
    if p.min == 0 && p.indxmin < p.indxmax{
        return p.max
    }
	if p.tempMin > 0 && p.indxmin > p.indxmax{
		return p.max - p.tempMin
	}
	

	if p.indxmin >= p.indxmax {
		return 0
	}
	return p.max - p.min
}

type ProductPrice struct {
	min, max, tempMin, indxmin, indxmax int
}