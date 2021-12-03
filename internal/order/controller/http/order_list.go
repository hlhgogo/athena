package http

import (
	"github.com/gin-gonic/gin"
	"github.com/hlhgogo/athena/pkg/extend"
)

// orderListReq 查询订单列表请求结构体
type orderListReq struct {
	OrderType uint8  `json:"orderType"` // 订单类型
	Page      uint64 `json:"page"`      // 页码
	PageSize  uint64 `json:"pageSize"`  // 分页数
}

// orderListResp 查询订单列表返回值
type orderListResp struct {
	ID   string `json:"id"`   // id
	Name string `json:"name"` // name
}

// OrderList 查询订单列表
func OrderList(ctx *gin.Context) {
	var r orderListReq
	if err := ctx.ShouldBind(&r); err != nil {
		extend.SendData(ctx, nil, err)
		return
	}

	panic(1)
	resp := orderListResp{
		ID:   "111",
		Name: "2222",
	}

	extend.SendSuccess(ctx, resp)
}
