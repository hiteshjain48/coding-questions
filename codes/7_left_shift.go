package codes

import "fmt"

func LeftShift(arr []int, k, length int) []int {
	shifted := make([]int, length)
	idx:=0
	for i:=0;i<length; i++{
		idx = (i + length - k)%length
		fmt.Print(idx)
		shifted[idx] = arr[i]
	}
	return shifted
}