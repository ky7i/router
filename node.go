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

func (n *node) addRouter(path string, handler http.HandlerFunc) {
	if n.path == "" && len(n.children) == 0 {
		child := &node{
			path:    path,
			indices: "",
			nType:   "static",
			handler: handler,
		}
		n.children = append(n.children, child)
	}

walk:
	for {
		i := longestCommonPrefix(path, n.path)

		if i < len(n.path) {
			child := &node{
				path:     n.path[i:],
				nType:    n.nType,
				children: n.children,
				handler:  n.handler,
			}
			n.path = n.path[:i]
			n.nType = "static"
			n.children = []*node{child} // check syntax
			n.handler = nil

			// sideeffect
			n.insertChild(path[i:], handler)
		}

		part := path[i:]
		for {
			j := 0
			if len(n.children) <= j {
				n.insertChild(part, handler)
				break walk
			}

			child := n.children[j]
			if child.path[0] == part[0] {
				n = child
				path = part // be carefull updating the wide scope variable
				continue walk
			}

			j++
		}

	}

}

func (n *node) insertChild(path string, handler http.HandlerFunc) {

}
