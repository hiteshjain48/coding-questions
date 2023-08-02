package occur

import "fmt"

func occurrences(arr int[], target int) int {
	length := len(arr)
	count := 0
	for i := 0; i < length; i++ {
		if arr[i] == target {
			count++	
		}
		if arr[i] > k {
			break
		}
	}
}
func main() {
	arr := {1,2,3,3,4,5,6,6,6,7}
	count = occurrences(arr, 6)
	print(count)
}