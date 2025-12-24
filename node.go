package router

import (
	"fmt"
	"net/http"
	"strings"
)

type node struct {
	// path is a segment of full path which is separated by common prefixes.
	path string

	// indices is the byte slice one of which is a first byte of a child node.
	// split any segments to segment[0] as indices and segment[0:] as a child node.
	//
	// ※invariant
	// len(n.children) == len(indices)
	// index of indices is connected with one of children, n.indices[i] was prefix of n.children[i]
	indices string

	children []*node

	nType   string
	handler http.HandlerFunc
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
	// n is root
	if n.path == "" && len(n.children) == 0 {
		child := &node{
			path:     path[0:],
			indices:  "",
			children: []*node{},
			nType:    "static",
			handler:  handler,
		}
		n.indices = string(path[0])
		n.children = append(n.children, child)
		return
	}

walk:
	// loop for , search in depth
	for {
		fmt.Printf("path: %q, node: %q\r\n", path, n.path)

		i := longestCommonPrefix(path, n.path)

		// split n
		if i < len(n.path) {
			child := &node{
				path:     n.path[i:],
				indices:  n.indices,
				nType:    n.nType,
				children: n.children,
				handler:  n.handler,
			}
			n.indices = string(n.path[i])
			n.path = n.path[:i]
			n.nType = "static"
			n.children = []*node{child} // check syntax
			n.handler = nil

			// side effect
			if i < len(path) {
				n.indices = n.indices + string(path[i])
				n.children = append(n.children, n.createChild(path[i:], handler))
			}
			return
		}

		if index := strings.Index(n.indices, string(path[i])); index == -1 {
			// TODO: indices should be created using append
			n.indices = n.indices + string(path[i])
			n.children = append(n.children, n.createChild(path[i:], handler))
			return
		} else {
			n = n.children[index]
			path = path[i:]
			continue walk
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
