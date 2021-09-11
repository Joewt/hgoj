package main

import (
	"time"

	"github.com/hnit-acm/hfunc/hapi"
	"github.com/hnit-acm/hfunc/hserver/hhttp"
	"github.com/yinrenxin/hgoj-api/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置，db
}

func Close() {

}

func main() {
	hapi.ServeTimeout("8080", nil, func(c *gin.Engine) {
		router.ApiRouter(c)
	}, Close, time.Second*5, hhttp.WithReadTimeout(time.Second*5), hhttp.WithReadTimeout(time.Second*10))
}
