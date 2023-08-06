package main

import (
	"fmt"
	"sort"
	"github.com/hiteshjain48/coding-questions/codes"
)

func main(){
	arr := [] int {1,2,3,4,5}
	sort.Ints(arr)
	k := 3
	fmt.Printf("k th largest number is %v\n", arr[k])
	// occur := codes.Occurrences(arr,k)
	count:= codes.PairsSum(arr,6)
	fmt.Printf("count is %d\n",count)
	max,min := codes.MinMax(arr)
	fmt.Printf("Min is %d and Max is %d\n", min,max)
	// fmt.Print(occur)
}