package api

import (
	"context"
	"fmt"
	"github.com/hlhgogo/gin-ext/extend"
	"github.com/hlhgogo/gin-ext/grequestsx"
	"github.com/hlhgogo/gin-ext/log"
	"github.com/levigross/grequests"
)

// DemoAPIReq ...
type DemoAPIReq struct {
	ID string `json:"id,omitempty"`
}

// DemoAPIResp ...
type DemoAPIResp struct {
	extend.Res
	Data []string `json:"data"`
}

func DemoAPI(ctx context.Context, req DemoAPIReq) error {
	ro := &grequests.RequestOptions{
		Context: ctx,
		JSON:    req,
	}

	url := "http://transocks-user-account.com/api/1/userinfo"
	resp, err := grequestsx.Post(url, ro, grequestsx.Flags{
		EnableLog:    true,
		DisableTrace: false,
	})
	if err != nil {
		return err
	}
	res := DemoAPIResp{}

	err = resp.JSON(&res)
	if err != nil {
		log.ErrorWithTrace(ctx, err, "get userinfo error")
		return err
	}

	if !res.Success {
		err := fmt.Errorf(res.Msg)
		log.ErrorWithTrace(ctx, err, "get userinfo error")
		return err
	}

	return nil
}
