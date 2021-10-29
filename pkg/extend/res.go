package extend

import (
	"github.com/gin-gonic/gin"
	athCtx "github.com/hlhgogo/athena/pkg/context"
	"github.com/hlhgogo/athena/pkg/errors"
	"net/http"
)

// Res api response结构
type Res struct {
	Success bool        `json:"success"`
	Code    int         `json:"code,omitempty"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceId,omitempty"`
}

// ResDebug 响应带debug信息
type ResDebug struct {
	*Res
	Debug map[string]interface{}
}

// SendSuccess 返回成功结果
func SendSuccess(ctx *gin.Context, resData interface{}) {
	var res *Res
	res = successRes()

	res.Data = resData
	res.TraceID = athCtx.GetTraceId(ctx)

	ctx.JSON(http.StatusOK, res)
}

// SendData 返回结果
// 根据err 是否为nil判断返回成功或失败
func SendData(ctx *gin.Context, resData interface{}, err error) {
	var res *Res
	var httpStatus = http.StatusOK
	if err != nil {
		res = failedRes()
		if e, ok := err.(*errors.Err); ok {
			httpStatus = http.StatusInternalServerError
			if code := e.Code(); code != 0 {
				res.Code = code
			}
			if msg := e.Message(); msg != "" {
				res.Msg = msg
			}
		} else if e, ok := err.(*errors.BadRequestError); ok {
			httpStatus = http.StatusBadRequest
			if code := e.Code(); code != 0 {
				res.Code = code
			}
			if msg := e.Message(); msg != "" {
				res.Msg = msg
			}
		} else if e, ok := err.(*errors.UnauthorizedError); ok {
			httpStatus = http.StatusUnauthorized
			if code := e.Code(); code != 0 {
				res.Code = code
			}
			if msg := e.Message(); msg != "" {
				res.Msg = msg
			}
		} else if e, ok := err.(*errors.ErrNotFoundError); ok {
			httpStatus = http.StatusNotFound
			if code := e.Code(); code != 0 {
				res.Code = code
			}
			if msg := e.Message(); msg != "" {
				res.Msg = msg
			}
		}
	} else {
		res = successRes()
	}

	res.Data = resData
	res.TraceID = athCtx.GetTraceId(ctx)

	ctx.JSON(httpStatus, res)
}

// newRes 新建
func newRes(success bool, code int, data interface{}) *Res {

	var msg string
	if value, ok := errors.ErrText[code]; ok {
		msg = value
	}
	return &Res{
		Success: success,
		Data:    data,
		Code:    code,
		Msg:     msg, // 原始，会被err中的msg替换 ，err中没有msg,会显示未定义
	}
}

// defaultRes 默认
func defaultRes() *Res {
	return newRes(true, errors.Success, struct{}{})
}

// failedRes 失败
func failedRes() *Res {
	return newRes(false, errors.ErrInternalServerError, struct{}{})
}

// successRes 成功
func successRes() *Res {
	return defaultRes()
}
