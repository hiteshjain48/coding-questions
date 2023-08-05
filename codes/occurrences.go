package codes

func Occurrences(arr []int, target int) int {
	length := len(arr)
	count := 0
	for i := 0; i < length; i++ {
		if arr[i] == target {
			count++	
		}
		if arr[i] > target {
			break
		}
	}
	return count
}
