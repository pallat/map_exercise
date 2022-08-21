package mapexercise

import "strings"

// TODO: split words and count them
func wordCount(s string) map[string]int {
	result := map[string]int{}
	elem := strings.Fields(s)
	for i := 0; i < len(elem); i++ {
		result[elem[i]]++
	}

	return result
}
