package leshy_elasticsearch_writer

import (
	"bytes"
	"context"
	"io"

	"github.com/google/uuid"
	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_elasticsearch "github.com/njxxdev/leshy/pkg/storages/elasticsearch"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type ElasticsearchWriter struct {
	name    string
	elastic *leshy_elasticsearch.Elasticsearch
	index   string
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
