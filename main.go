package main

import (
	"fmt"
	"sort"
	"github.com/hiteshjain48/coding-questions/codes"
)

func main(){
	arr := [] int {1,2,3,4,5}
	sort.Ints(arr)
	// k := 3
	// fmt.Printf("k th largest number is %v\n", arr[k])
	// // occur := codes.Occurrences(arr,k)
	// count:= codes.PairsSum(arr,6)
	// fmt.Printf("count is %d\n",count)
	// max,min := codes.MinMax(arr)
	// fmt.Printf("Min is %d and Max is %d\n", min,max)
	// // fmt.Print(occur)
	// shifted:= codes.LeftShift(arr,2,len(arr))
	// fmt.Println(shifted)
	// anagram := codes.Anagram("listen","hitesh")

	// fmt.Println(anagram)
	// var reverse string
	// str := "Vaibhya bsdk"
	// length := len(str)
	// codes.ReverseRecursion(str,&reverse,length)
	// fmt.Println(reverse)

	// char := codes.FirstNotRepeating("programming")
	// fmt.Println(char)
	// codes.ReverseSentence("I love programming")
	balanced := codes.BalancedString("(){}")
	fmt.Println(balanced)
}