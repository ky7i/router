package router

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	list := []string{
		"/user",
		"/users",
		"/user/",
		"/us",
		"/usx",
		"/",
	}
	path := "/user"

	expected := []int{5, 5, 4, 3, 1}

	for j := range list {
		if i := longestCommonPrefix(path, list[j]); i != expected[j] {
			t.Errorf("got %d, want %d", i, expected[j])
		}
	}
}
