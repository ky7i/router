package router

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestAddRouter(t *testing.T) {

	n := &node{
		path:     "",
		children: []*node{},
		handler:  nil,
	}

	dummyHandler := func(_ http.ResponseWriter, _ *http.Request) {}

	n.addRouter("/user", dummyHandler)
	n.addRouter("/user/userId", dummyHandler)
	n.addRouter("/user/profile", dummyHandler)
	n.addRouter("/us", dummyHandler)
	n.addRouter("/apis/parking", dummyHandler)
	n.addRouter("/usa", dummyHandler)

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
