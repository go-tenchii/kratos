package http

import "github.com/gorilla/mux"

// HandleOption is route option.
type HandleOption func(*handleOptions)

type handleOptions struct {
	methods []string
}

// Methods with handle methods.
func Methods(methods ...string) HandleOption {
	return func(o *handleOptions) {
		o.methods = methods
	}
}

func applyHandleOpts(route *mux.Route, opts ...HandleOption) {
	options := handleOptions{}
	for _, o := range opts {
		o(&options)
	}
	if len(options.methods) > 0 {
		route.Methods(options.methods...)
	}
}
