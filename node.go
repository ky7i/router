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
		return
	}

walk:
	for {
		// go next loop when root
		if n.path == "" {
			n = n.children[0]
			continue walk
		}

		i := longestCommonPrefix(path, n.path)

		// split n
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

			// side effect
			child = n.createChild(path[i-1:], handler)
			n.children = append(n.children, child)
			return
		}

		// make new nodes
		part := path[i-1:]
		for {
			j := 0
			if len(n.children) <= j {
				child := n.createChild(part, handler)
				n.children = append(n.children, child)
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

func (n *node) createChild(path string, handler http.HandlerFunc) *node {
	node := &node{
		path:     path,
		children: []*node{},
		handler:  handler,
	}
	return node
}
