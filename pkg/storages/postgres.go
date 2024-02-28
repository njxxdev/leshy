package storages

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/njxxdev/leshy/pkg/component"
	"github.com/njxxdev/leshy/pkg/config"
)

type PostgresRepository struct {
	name string
	url  string
	Pool *pgxpool.Pool
}

func (comp PostgresRepository) GetInstance() component.Component {
	return comp
}

func (comp PostgresRepository) GetName() string { return comp.name }

func NewPostgresRepository(name string) *PostgresRepository {
	url := config.GetConfigs().GetParameters()[name].(map[string]interface{})["url"].(string)
	pool, err := pgxpool.Connect(context.Background(), url)

	if err != nil {
		panic("Ошибка при подключении к базе данных:" + err.Error())
	}

	return &PostgresRepository{
		name: name,
		url:  url,
		Pool: pool,
	}
}
