package str_test

import (
	"testing"

	"github.com/Drew-Kimberly/go-exploration/str"
)

var testCases = []struct {
	haystack string
	needle   string
	expected int
}{
	{"hello", "el", 1},
	{"mississippi", "issippi", 4},
	{"aaa", "a", 0},
	{"abc", "abc", 0},
	{"test", "", 0},
	{"abc", "d", -1},
	{"aaaab", "aaaabc", -1},
}

func TestIndexOf(t *testing.T) {
	for _, tc := range testCases {
		result := str.IndexOf(tc.haystack, tc.needle)
		if result != tc.expected {
			t.Errorf("TestIndexOf failure: IndexOf('%s', '%s'), got %d, expected %d", tc.haystack, tc.needle, result, tc.expected)
		}
	}
}
