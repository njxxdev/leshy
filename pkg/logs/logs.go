package leshy_logs

import (
	"log/slog"
	"os"

	leshy_component "github.com/njxxdev/leshy/pkg/component"
)

type Logger struct {
	name string
	log  *slog.Logger
}

func (comp *Logger) Instance() leshy_component.Component {
	return comp
}

func (comp *Logger) Name() string { return comp.name }

func NewAPIServer(name string) *Logger {

	return &Logger{
		name: name,
		log:  slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
	}
}
