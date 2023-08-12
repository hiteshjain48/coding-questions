package codes

import "fmt"

func ReverseSentence(str string) {
	var start, end int
	start = len(str)-1
	end = len(str)-1
	for i:=len(str)-1;i>=0;i--{
		if string(str[i])==" " || i == 0{
			end = i-1
			if i==0 {
				end = 0
			}
			reverse := str[end:start+1]
			fmt.Printf("%v",reverse)
			start = end+1
		}
	}
	fmt.Println()
}