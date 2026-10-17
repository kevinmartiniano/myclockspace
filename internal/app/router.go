package app

import (
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	r.handle(http.MethodGet, path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	r.handle(http.MethodPost, path, handler)
}

func (r *Router) Put(path string, handler http.HandlerFunc) {
	r.handle(http.MethodPut, path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) {
	r.handle(http.MethodDelete, path, handler)
}

func (r *Router) Patch(path string, handler http.HandlerFunc) {
	r.handle(http.MethodPatch, path, handler)
}

func (r *Router) handle(method string, path string, handler http.HandlerFunc) {
	if !allowedMethods(method) {
		return
	}

	r.mux.HandleFunc(method+" "+path, handler)
}

var allowedMethodsMap = map[string]struct{}{
	http.MethodGet:    {},
	http.MethodPost:   {},
	http.MethodPut:    {},
	http.MethodDelete: {},
	http.MethodPatch:  {},
}

func allowedMethods(method string) bool {
	_, exists := allowedMethodsMap[method]
	return exists
}
