package leshy_postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
)

type PostgresRepository struct {
	name string
	url  string
	Pool *pgxpool.Pool
}

func (comp PostgresRepository) Instance() leshy_component.Component {
	return comp
}

func (comp PostgresRepository) Name() string { return comp.name }

func New(name string) *PostgresRepository {
	url := leshy_config.Get().Parameters()[name].(map[string]interface{})["url"].(string)
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
