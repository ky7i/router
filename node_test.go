package router

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetValue(t *testing.T) {
	// n := createRouter()

}

var testStr string

func fakeHandler(str string) func(http.ResponseWriter, *http.Request) {
	return func(http.ResponseWriter, *http.Request) {
		testStr = str
	}
}

func createRouter() *node {
	n := &node{
		path:     "",
		children: []*node{},
		handler:  nil,
	}

	paths := []string{
		"/user",
		"/user/userId",
		"/user/profile",
		"/us",
		"/apis/parking",
		"/usa",
	}

	for _, path := range paths {
		n.addRouter(path, fakeHandler(path))
	}
	return n
}

func TestAddRouter(t *testing.T) {
	n := createRouter()
	n.walk(0)
}

func (n *node) walk(depth int) {
	if n == nil {
		return
	}

	fmt.Println(strings.Repeat("  ", depth), n.path)

	for _, child := range n.children {
		child.walk(depth + 1)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	list := []string{
		"/user",
		"/users",
		"/user/",
		"/us",
		"/usx",
		"/",
		"",
	}
	path := "/user"

	expected := []int{5, 5, 5, 3, 3, 1}

	for j := range list {
		if i := longestCommonPrefix(path, list[j]); i != expected[j] {
			t.Errorf("got %d, want %d", i, expected[j])
		}
	}
}
