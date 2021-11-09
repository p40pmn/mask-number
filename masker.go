package masknumber

func Masker(s string) string {
	var str string
	if len(s) < 8 {
		return s
	}

	for i := 0; i < len(s); i++ {
		if i <= len(s)-3 && i >= len(s)-8 {
			str += "X"
		} else {
			str += string(s[i])
		}
	}
	return str
}
