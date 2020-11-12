package client

import "net/http"

// ServeMux is a wrapper around http.ServeMux that instruments handlers for tracing.
type ServeMux struct {
	mux *http.ServeMux
}

// NewServeMux creates a new TracedServeMux.
func NewServeMux() *ServeMux {
	return &ServeMux{
		mux: http.NewServeMux(),
	}
}

// Handle implements http.ServeMux#Handle
func (sm *ServeMux) Handle(pattern string, handler http.Handler) {
	sm.mux.Handle(pattern, handler)
}

// ServeHTTP implements http.ServeMux#ServeHTTP
func (sm *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sm.mux.ServeHTTP(w, r)
}
