package codes

func PairsSum(arr []int, sum int) int {
	count := 0
	for i:=0;i<len(arr)-1;i++{
		for j:=i+1;j<len(arr);j++{
			if arr[i]+arr[j]==sum{
				count++
			}
		}
	}
	return count
}