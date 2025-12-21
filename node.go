package router

import (
	"net/http"
	"strings"
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
	// loop for , search in depth
	for {
		// go to the next loop when root
		// TODO refactor
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
			n.insertChild(path[i-1:], handler)
			return
		}

		// make new nodes
		part := path[i-1:]
		// loop for children, search in breadth
		// 		for j := range n.children {
		// 			if len(n.children) <= j {
		// 				n.insertChild(part, handler)
		// 				break walk
		// 			}
		//
		// 			child := n.children[j]
		// 			if child.path[0] == part[0] {
		// 				n = child
		// 				path = part // be carefull updating the wide scope variable
		// 				continue walk
		// 			}
		// 		}

		if index := strings.Index(n.indices, string(part[0])); index == -1 {
			n.insertChild(part, handler)
			return
		} else {
			n = n.children[index]
			continue walk
		}
	}
}

func (n *node) insertChild(path string, handler http.HandlerFunc) {
	node := &node{
		path:     path,
		children: []*node{},
		handler:  handler,
	}
	n.children = append(n.children, node)
	n.indices = n.indices + string(path[0])
}
