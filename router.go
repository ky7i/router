package router

import (
	"log"
	"net/http"
)

type Router struct {
	// separate trees by HTTP methods.
	// ex) trees[0] is Tree of GET method.
	trees [10]*node

	NotFound http.HandlerFunc
}

func New() *Router {
	r := &Router{trees: [10]*node{}}
	for i := range r.trees {
		r.trees[i] = &node{}
	}
	r.NotFound = http.NotFound
	return r
}

func (r *Router) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.insert("GET", path, handler)
}

func (r *Router) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	r.insert("POST", path, handler)
}

func (r *Router) insert(method string, path string, handler http.HandlerFunc) {
	if path == "" {
		panic("Registering path must not be empty.")
	} else if path[0] != '/' {
		panic("Path must have the prefix '/'.")
	}

	switch method {
	case "GET":
		r.trees[0].addRouter(path, handler)
	case "POST":
		r.trees[1].addRouter(path, handler)
	}
}

// TODO: write more pattern, detail.
func getMethodIndexOf(method string) int {
	switch method {
	case "GET":
		return 0
	case "POST":
		return 1
	}
	return -1
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	log.Printf("method: %q, path: %q\r\n", method, path)
	methodIndex := getMethodIndexOf(method)
	if node := r.trees[methodIndex].getValue(path); node != nil {
		node.handler(w, req)
		return
	}
	r.NotFound(w, req)
}
