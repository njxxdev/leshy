package leshy_http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
)

type APIServer struct {
	name   string
	engine *gin.Engine
}

func (comp *APIServer) Instance() leshy_component.Component {
	return comp
}

func (comp *APIServer) Name() string { return comp.name }

func NewAPIServer(name string, engine *gin.Engine) *APIServer {
	engineE := engine
	if engineE == nil {
		engineE = gin.Default()
	}
	return &APIServer{
		name:   name,
		engine: engineE,
	}
}

func (serv *APIServer) AddHandlers(handlers ...Handler) *APIServer {
	for _, h := range handlers {
		if h.process == nil {
			panic("API: Handler \"" + h.method + " " + h.path + "\" has no func to process context")
		}
		serv.engine.Handle(h.method, h.path, h.process)
	}
	return serv
}

func (serv *APIServer) Run() error {
	port := leshy_config.Get().Parameters()[serv.name].(map[string]interface{})["port"].(int)
	return serv.engine.Run(":" + strconv.Itoa(port))
}
