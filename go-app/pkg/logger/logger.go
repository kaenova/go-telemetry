package logger

import (
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

type LoggerOptions struct {
	AppName string

	ElasticHost string
	ElasticPort string
}

type Logger struct {
	Config LoggerOptions
	Log    *logrus.Logger
}

var DefaultLoggerOption LoggerOptions = LoggerOptions{
	AppName: "hello-world",
}

func NewLogger(opt LoggerOptions) *Logger {

	log := logrus.New()

	if opt.AppName == "" {
		opt.AppName = "hello-world"
	}

	if opt.ElasticHost != "" && opt.ElasticPort != "" {
		elasticEndpoint := fmt.Sprintf("http://%s:%s", opt.ElasticHost, opt.ElasticPort)
		client, err := elastic.NewClient(elastic.SetURL(elasticEndpoint))
		if err != nil {
			log.Panic(err)
		}
		hook, err := elogrus.NewAsyncElasticHook(client, opt.ElasticHost, logrus.DebugLevel, opt.AppName)
		if err != nil {
			log.Panic(err)
		}
		log.Hooks.Add(hook)
	}

	return &Logger{
		Log:    log,
		Config: opt,
	}
}
