package codes

type Stack struct {
	stack []string
}

func (s Stack) push(value string) {
	s.stack = append(s.stack, value)
}

func (s Stack) pop() {
	if len(s.stack)!=0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func BalancedString(str string) bool{
	var balanced bool
	s := Stack{}
	for i:=0; i<len(str);i++ {
		if string(str[i]) == "(" || string(str[i]) == "[" || string(str[i]) == "{"{
			s.push(string(str[i]))
		} else if string(str[i]) == ")" && string(s.stack[len(s.stack)-1]) == "(" || string(str[i]) == "]" && string(s.stack[len(s.stack)-1]) == "[" || string(str[i]) == "}" && string(s.stack[len(s.stack)-1]) == "{"{
			s.pop()
		}
	}
	if len(s.stack) == 0 {
		balanced = true
	}
	return balanced
}