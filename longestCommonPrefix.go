package main
import(
	"fmt"
)

func main(){
	s := []string{"av","a"}
	fmt.Println(longestCommonPrefix(s))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1{
		return strs[0]
	}
    res := ""
	minLenWord := len(strs[0])
	for _, v := range strs{		
		if len(v) < minLenWord{
			minLenWord = len(v)
		}
		if len(v) == 0{
			return ""
		}
	}
	i := 0
	for  {		
		tmp := ""
		for j := 0; j < len(strs)-1; j++ {
			if i + 1 > minLenWord{
				res = res + tmp
				return res
			}
			if strs[j][i] == strs[j+1][i]{
				if tmp == string(strs[j][i]){
					continue
				}				
				tmp = tmp + string(strs[j][i])
			}else{
				return res
			}
			
			
			}
			res = res + tmp
			i++
		}
	}


