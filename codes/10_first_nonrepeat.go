package codes

// import "fmt"



func FirstNotRepeating(str string) string {
	hashMap := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		hashMap[str[i]] += 1	
	}
	// fmt.Println(hashMap)
	for key, value := range hashMap{
		if value == 1 {
			return string(key)
		}
	}
	return "-1"
}