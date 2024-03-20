package leshy_elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v7"

	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
)

type Elasticsearch struct {
	name   string
	client *elasticsearch.Client
}

// Instance - возвращает экземпляр компонента
func (comp *Elasticsearch) Instance() leshy_component.Component {
	return comp
}

// Name - возвращает имя компонента
func (comp *Elasticsearch) Name() string {
	return comp.name
}

func (comp *Elasticsearch) Client() *elasticsearch.Client {
	return comp.client
}

// New - создает новый экземпляр компонента
func New(name string) *Elasticsearch {
	config := leshy_config.Get().Parameters()[name].(map[string]interface{})

	// Подключение к Elasticsearch
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config["addresses"].([]string),
		Username:  config["username"].(string),
		Password:  config["password"].(string),
	})

	if err != nil {
		panic("Elasticsearch: " + err.Error())
	}

	return &Elasticsearch{
		name:   name,
		client: client,
	}
}
