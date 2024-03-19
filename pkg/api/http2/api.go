package leshy_http2

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
	path    string
	methods map[string]func(*gin.Context)
}

type APIServer struct {
	name   string
	engine *gin.Engine

	handlers []handler
}

func (server *APIServer) Instance() interface{} {
	return server
}

func (server *APIServer) Name() string { return server.name }

func (server *APIServer) updateEngine(handlers []handler) {
	for _, h := range handlers {
		for method, process := range h.methods {
			server.engine.Handle(method, h.path, process)
		}
	}
}

func New(name string, engine *gin.Engine) *APIServer {
	engineE := engine
	if engineE == nil {
		engineE = gin.Default()
	}
	return &APIServer{
		name:   name,
		engine: engineE,
	}
}

func (server *APIServer) Append(
	paths []string,
	methods map[string]func(*gin.Context),
) *APIServer {
	handlers := []handler{}
	for _, path := range paths {
		handlers = append(handlers,
			handler{
				path:    path,
				methods: methods,
			},
		)
	}
	server.handlers = append(server.handlers, handlers...)
	server.updateEngine(handlers)
	return server
}
