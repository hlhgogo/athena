package event

import (
	"github.com/hlhgogo/athena/application/event/subscribe"
	"github.com/hlhgogo/eventbus"
	"github.com/hlhgogo/gin-ext/log"
)

// EventBusHandler ...
var EventBusHandler eventbus.Bus

// Init 初始化
func Init() {
	log.Info("init eventbus.")
	EventBusHandler = eventbus.New()
	EventBusHandler.SubscribeAsync(subscribe.NewDemoSubscribe(), false)
}

// Quit Quit
func Quit() {
	log.Info("eventbus exit.")
	EventBusHandler.Quit()
}
