package leshy_minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
)

type MinIOComponent struct {
	name string

	Client *minio.Client
}

func (comp MinIOComponent) Instance() leshy_component.Component {
	return comp
}

func (comp MinIOComponent) Name() string { return comp.name }

// NewMinIOComponent - создание нового компонента MinIO
func NewMinIOComponent(name string) *MinIOComponent {
	endpoint := leshy_config.GetConfigs().GetParameters()[name].(map[string]interface{})["endpoint"].(string)
	accessKeyID := leshy_config.GetConfigs().GetParameters()[name].(map[string]interface{})["accessKeyID"].(string)
	secretAccessKey := leshy_config.GetConfigs().GetParameters()[name].(map[string]interface{})["secretAccessKey"].(string)
	useSSL := leshy_config.GetConfigs().GetParameters()[name].(map[string]interface{})["useSSL"].(bool)

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
