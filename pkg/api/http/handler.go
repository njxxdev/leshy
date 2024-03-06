package leshy_http

import "github.com/gin-gonic/gin"

type Handler struct {
	method string
	path   string

	auth    func(*gin.Context) bool
	process func(*gin.Context)
}

func NewHandler(
	method string,
	path string,
	process func(*gin.Context),
	auth func(*gin.Context) bool,
) *Handler {
	return &Handler{
		method:  method,
		path:    path,
		auth:    auth,
		process: process,
	}
}

func NewHandlerExtended(
	path string,
	processMap map[string]func(*gin.Context),
	auth func(*gin.Context) bool,
) []Handler {
	h := []Handler{}
	for method, process := range processMap {
		h = append(h, Handler{
			method:  method,
			path:    path,
			auth:    auth,
			process: process,
		})
	}
	return h
}

func NewMultipathHandler(
	method string,
	paths []string,
	process func(*gin.Context),
	auth func(*gin.Context) bool,
) []Handler {
	handlers := []Handler{}
	for _, path := range paths {
		handlers = append(handlers,
			Handler{
				method:  method,
				path:    path,
				auth:    auth,
				process: process,
			})
	}
	return handlers
}

func NewMultipathHandlerExtended(

	paths []string,
	processMap map[string]func(*gin.Context),
	auth func(*gin.Context) bool,
) []Handler {
	handlers := []Handler{}
	for method, process := range processMap {
		for _, path := range paths {
			handlers = append(handlers,
				Handler{
					method:  method,
					path:    path,
					auth:    auth,
					process: process,
				})
		}
	}
	return handlers
}
