package module01

import (
	"fmt"
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//	BaseToDec("E", 16) => 14
//	BaseToDec("1110", 2) => 14
func BaseToDec(value string, base int) int {

	const charset = "0123456789ABCDEF"

	var result int

	fmt.Println(value)
	// use Reverse() from the same package!
	for i, val := range Reverse(value) {
		v := strings.Index(charset, string(val))
		result += int(math.Pow(float64(base), float64(i))) * v
	}
	return result
}

// should there be no Reverse(), then just use the index and the length of string:
// power := (len(value) - 1) - i
// result += int(math.Pow(float64(base), float64(power))) * v
