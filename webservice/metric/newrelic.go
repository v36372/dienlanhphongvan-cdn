package metric

import (
	"time"

	"github.com/newrelic/go-agent"
)

type newRelic struct {
	newrelic.Application
	Enable bool
}

var newrelicApp newRelic

func NewRelic() newRelic {
	return newrelicApp
}

func InitNewrelic(enable bool, name, license string) {
	newrelicApp.Enable = enable
	if !enable {
		return
	}
	cfg := newrelic.NewConfig(name, license)
	nr, err := newrelic.NewApplication(cfg)
	if err != nil {
		panic(err)
	}

	newrelicApp.Application = nr
}

func CloseNewrelic() {
	if newrelicApp.Application == nil {
		return
	}
	newrelicApp.Shutdown(5 * time.Second)
}
