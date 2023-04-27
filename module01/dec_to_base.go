package module01

import "fmt"

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {

	if base > 16 {
		panic("The function handles max base 16")
	}

	var answer string
	for dec > 0 {
		digit := dec % base
		answer = intToHex(digit) + answer
		dec = dec / base
	}

	return answer
}

func intToHex(i int) string {
	if i > 9 {
		var result string
		switch {
		case i == 10:
			result = "A"
		case i == 11:
			result = "B"
		case i == 12:
			result = "C"
		case i == 13:
			result = "D"
		case i == 14:
			result = "E"
		case i == 15:
			result = "F"
		}
		return result
	}
	return fmt.Sprint(i)
}
