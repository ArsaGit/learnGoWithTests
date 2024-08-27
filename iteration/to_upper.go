package iteration

func MyToUpper(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		result += string(s[i] - 32)
	}
	return result
}
