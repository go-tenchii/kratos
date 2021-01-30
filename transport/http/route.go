package http

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

// RouteGroup .
type RouteGroup struct {
	root   string
	router *mux.Router
}

func (r *RouteGroup) path(p string) string {
	return path.Join(r.root, p)
}

// GET .
func (r *RouteGroup) GET(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("GET")
}

// HEAD .
func (r *RouteGroup) HEAD(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("HEAD")
}

// POST .
func (r *RouteGroup) POST(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("POST")
}

// PUT .
func (r *RouteGroup) PUT(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("PUT")
}

// DELETE .
func (r *RouteGroup) DELETE(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("DELETE")
}

// PATCH .
func (r *RouteGroup) PATCH(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("PATCH")
}

// OPTIONS .
func (r *RouteGroup) OPTIONS(path string, h http.HandlerFunc) {
	r.router.HandleFunc(r.path(path), h).Methods("OPTIONS")
}
