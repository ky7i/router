package router

import (
	"net/http"
)

type node struct {
	path     string
	indices  string
	nType    string
	children []*node
	handler  http.HandlerFunc
}

func longestCommonPrefix(a, b string) int {
	i := 0
	max := min(len(a), len(b))
	for i < max && a[i] == b[i] {
		i++
	}
	return i
}

func (n *node) addRouter(path string, handler Handler) {
	if n.path == "" && len(n.children) == 0 {
		child := &node{
			path: path,
			indices: "",
			nType: "static",
			handler: handler,
		}
		n.children = append(n.children, child)
	}

	for {
		i := longestCommonPrefix(path, n.path)
	}

}

func (n *node) insertChild(path, fullpath string, handler Handler) {

}
