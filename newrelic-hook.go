package logrus_newrelic_hook

import (
	"bytes"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type NewrelicHook struct {
	apiKey string
}

const BaseUrl = "https://log-api.newrelic.com/log/v1"

func NewNewrelicHook(apiKey string) *NewrelicHook {
	return &NewrelicHook{
		apiKey: apiKey,
	}
}

func (hook *NewrelicHook) Fire(entry *logrus.Entry) error {
	var url = BaseUrl + "?Api-Key=" + hook.apiKey
	var jsonData, err = entry.Bytes()
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (hook *NewrelicHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
