package module01

import (
	"fmt"
	"strconv"
)

// BaseToBase takes in a number, the base it is currently
// in, and the base you want it to be converted to. It then
// returns a string representing the number in the new base.
//
// Eg:
//
//	BaseToBase("E", 16, 2) => "1110"
func BaseToBase(value string, base, newBase int) string {
	if base == newBase {
		fmt.Printf("base: %v egauls to newBase: %v", base, newBase)
		return value
	} else if base == 10 {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("can't convert %v to integer", err)
			// return
		}
		return DecToBase(intValue, newBase)
	} else if newBase == 10 {
		return strconv.Itoa(BaseToDec(value, base))
	}
	// this would do the task itself, but the above conditionals improve performance
	// because only the necessary functions are called (see benchmark in tests)
	return DecToBase(BaseToDec(value, base), newBase)
}
