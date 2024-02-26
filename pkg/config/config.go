package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type configManager struct {
	filename string
	config   map[interface{}]interface{}
}

var instance *configManager
var once sync.Once

func LoadConfigs(filename string) {
	once.Do(func() {
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		var result map[interface{}]interface{}
		err = yaml.Unmarshal(data, &result)
		if err != nil {
			panic(err)
		}
		instance = &configManager{
			filename: filename,
			config:   result,
		}
	})
}
func GetConfigs() *configManager {
	if instance == nil {
		panic("Configs: Load configs berfore use")
	}
	return instance
}

func (config *configManager) GetParameters() map[interface{}]interface{} {
	return GetConfigs().config
}
