package leshy_elasticsearch_writer

import (
	"bytes"
	"context"
	"io"

	"github.com/google/uuid"
	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
	leshy_elasticsearch "github.com/njxxdev/leshy/pkg/storages/elasticsearch"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchWriter struct {
	name    string
	elastic *leshy_elasticsearch.Elasticsearch
	index   string
}

func New(name string) *ElasticsearchWriter {
	config := leshy_config.Get().Parameters()[name].(map[string]interface{})

	elastic_name := config["source"].(string)
	elastic_comp, err := leshy_component.GetContext().GetComponent(elastic_name)
	if err != nil {
		panic("ElasticsearchWriter: " + err.Error())
	}

	return &ElasticsearchWriter{
		name:    name,
		elastic: elastic_comp.(*leshy_elasticsearch.Elasticsearch),
		index:   config["index"].(string),
	}
}

func (writer *ElasticsearchWriter) Instance() leshy_component.Component {
	return writer
}

func (writer *ElasticsearchWriter) Name() string {
	return writer.name
}

func (writer *ElasticsearchWriter) Write(p []byte) (n int, err error) {
	req := esapi.IndexRequest{
		Index:      writer.index,
		DocumentID: uuid.NewString(),
		Refresh:    "true",
		Body:       bytes.NewReader(p),
	}

	res, err := req.Do(context.Background(), writer.elastic.Client())
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return 0, io.EOF
	}

	return 0, nil
}
