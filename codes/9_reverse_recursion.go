package codes


func ReverseRecursion(str string, reverse *string, length int){
	if length == 0 {
		return
	}
	*reverse = *reverse + string(str[length-1])
	length--
	ReverseRecursion(str[:length], reverse,length)
}