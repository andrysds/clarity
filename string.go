package clarity

import "strconv"

func Atoi(number string) int {
	result, err := strconv.Atoi(number)
	if err != nil {
		return 0
	}
	return result
}
