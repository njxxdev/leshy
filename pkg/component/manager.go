package component

import (
	"errors"
	"reflect"
	"sync"
)

type componentManager struct {
	components map[string]Component
}

var instance *componentManager
var once sync.Once

// Функция, возвращающая экземпляр Singleton, реализует интерфейс Singleton
func GetComponentManager() *componentManager {
	once.Do(func() {
		instance = &componentManager{
			components: make(map[string]Component),
		}
	})
	return instance
}

func (manager *componentManager) Append(components ...Component) *componentManager {
	countComponents := len(components)
	if countComponents > 0 {
		for i := 0; i < countComponents; i++ {
			value, ok := GetComponentManager().components[components[i].GetName()]
			if ok && reflect.TypeOf(value) != reflect.TypeOf(components[i]) {
				panic("ComponentManager: Reinitialization component with same name \"" + components[i].GetName() + "\" by different type")
			}
			GetComponentManager().components[components[i].GetName()] = components[i]
		}
	}
	return manager
}

func (manager *componentManager) GetComponent(name string) (Component, error) {
	value, ok := GetComponentManager().components[name]
	if ok {
		return value, nil
	}
	return nil, errors.New("not found component: " + name)
}
