package router

import (
	"net/http"
)

type Router struct {
	tree []*node
}

func (r *Router) GET(path string, handler http.HandlerFunc) {

}

func (r *Router) POST(path string, handler http.HandlerFunc) {

}

func (r *Router) ServeHTTP(w http.ResponseWriter, h *http.Request) {

}
