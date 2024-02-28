package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/njxxdev/leshy/pkg/component"
	"github.com/njxxdev/leshy/pkg/config"
)

type APIServer struct {
	name   string
	engine *gin.Engine
}

func (comp *APIServer) GetInstance() component.Component {
	return comp
}

func (comp *APIServer) GetName() string { return comp.name }

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
	port := config.GetConfigs().GetParameters()[serv.name].(map[string]interface{})["port"].(int)
	return serv.engine.Run(":" + strconv.Itoa(port))
}
