package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/njxxdev/leshy/pkg/component"
	"github.com/njxxdev/leshy/pkg/config"
)

type MinIOComponent struct {
	name string

	Client *minio.Client
}

func (comp MinIOComponent) GetInstance() component.Component {
	return comp
}

func (comp MinIOComponent) GetName() string { return comp.name }

// NewMinIOComponent - создание нового компонента MinIO
func NewMinIOComponent(name string) *MinIOComponent {
	endpoint := config.GetConfigs().GetParameters()[name].(map[string]interface{})["endpoint"].(string)
	accessKeyID := config.GetConfigs().GetParameters()[name].(map[string]interface{})["accessKeyID"].(string)
	secretAccessKey := config.GetConfigs().GetParameters()[name].(map[string]interface{})["secretAccessKey"].(string)
	useSSL := config.GetConfigs().GetParameters()[name].(map[string]interface{})["useSSL"].(bool)

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic("Ошибка при подключении к MinIO:" + err.Error())
	}

	return &MinIOComponent{
		name:   name,
		Client: client,
	}
}
