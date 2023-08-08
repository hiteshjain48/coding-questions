package codes

func Anagram(str1, str2 string) bool{
	hashMap := make(map[byte]int)
	if len(str1) != len(str2) {
		return false
	}
	for i:=0;i<len(str1);i++{
		hashMap[str1[i]] += 1
		hashMap[str2[i]] -= 1
	}
	for _, element := range hashMap {
		if element != 0 {
			return false
		}
	}
	return true

}