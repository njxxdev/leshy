package leshy_logs

import (
	"io"
	"log/slog"
	"os"

	leshy_component "github.com/njxxdev/leshy/pkg/component"
	leshy_config "github.com/njxxdev/leshy/pkg/config"
)

type Logger struct {
	name string
	log  *slog.Logger
}

func (comp *Logger) Instance() leshy_component.Component {
	return comp
}

func (comp *Logger) Name() string { return comp.name }

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
		elastic_component_name := config["component"].(string)
		elastic, err := leshy_component.GetContext().GetComponent(elastic_component_name)
		if err != nil {
			panic("Logger: " + err.Error())
		}
		_ = elastic
		// writer = elastic.(*Elasticsearch).log
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
		log: slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
			Level: logLevel,
		})),
	}
}
