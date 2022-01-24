package subscribe

import (
	"context"
	"github.com/hlhgogo/gin-ext/log"
)

// demoSubscribe 订阅示例
type demoSubscribe struct {
	event interface{} `subscribe:"RunEvent" topic:"demoTopic" quit:"Quit"`
}

// NewDemoSubscribe 更新订单状态
func NewDemoSubscribe() *demoSubscribe {
	return new(demoSubscribe)
}

// RunEvent 执行
func (event *demoSubscribe) RunEvent(ctx context.Context, id string) {
	log.InfoWithTrace(ctx, "DemoSubscribe RunEvent Params:%s", id)
}

// Quit 退出
func (event *demoSubscribe) Quit() {

}
