package leshy_logs

import (
	"io"
	"log/slog"
	"os"

	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
	leshy_elasticsearch_writer "github.com/njxxdev/leshy/pkg/storages/elasticsearch/writer"
)

type Logger struct {
	name string
	log  *slog.Logger
}

func (comp *Logger) Instance() leshy_component.Component {
	return comp
}

func (comp *Logger) Name() string { return comp.name }

func (comp *Logger) Log() *slog.Logger { return comp.log }

func New(name string) *Logger {
	config := leshy_config.Get().Parameters()[name].(map[string]interface{})

	// dest: stdout, stderr, file, elasticsearch
	dest := config["dest"].(string)
	var writer io.Writer
	if dest == "stdout" {
		writer = os.Stdout
	} else if dest == "stderr" {
		writer = os.Stderr
	} else if dest == "file" {
		file, err := os.OpenFile(
			config["file"].(string),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			panic("Logger: " + err.Error())
		}
		writer = file
	} else if dest == "elasticsearch" {
		elastic_writer_name := config["component"].(string)
		elastic_writer_comp, err := leshy_component.GetContext().GetComponent(elastic_writer_name)
		if err != nil {
			panic("Logger: " + err.Error())
		}
		writer = elastic_writer_comp.(*leshy_elasticsearch_writer.ElasticsearchWriter)
	} else {
		panic("Logger: Unknown destination \"" + dest + "\"")
	}

	// Log level: debug, info, warn, error, fatal, panic
	level := config["level"].(string)
	var logLevel slog.Level
	switch level {

	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		panic("Logger: Unknown log level \"" + level + "\"")
	}

	return &Logger{
		name: name,
		log: slog.New(slog.NewJSONHandler(writer,
			&slog.HandlerOptions{
				Level: logLevel,
			},
		)),
	}
}
