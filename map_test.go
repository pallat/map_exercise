package mapexercise

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	given := "Apple Banana Apple Banana apple"
	want := map[string]int{
		"Apple":  2,
		"Banana": 2,
		"apple":  1,
	}

	result := wordCount(given)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("wordCount(%s) = %v, want %v", given, result, want)
	}
}
