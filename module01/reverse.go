package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	var result string
	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}
	return result
}

func ReverseStringBuilder(word string) string {
	var result strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		result.WriteByte(word[i])
	}
	return result.String()
}

// this works with non ASCII characters (uses runes instead of bytes),
// also, it build the result from the first character of the word!
func ReverseWithRunes(word string) string {
	var result string
	for _, r := range word {
		result = string(r) + result
	}
	return result
}
