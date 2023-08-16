package codes


func GreatestCommonDivisor(a,b int) int {
	remainder := 1
	var ans int
	if b > a {
		temp := a
		a = b
		b = temp
	}
	if a%b == 0 {
		return b
	}
	for remainder > 0{
		ans = remainder
		remainder = a%b
		a = b
		b = remainder
	}
	return ans
}