package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yinrenxin/hgoj-api/internal/types"
)

// TODO: wire解决依赖
var User = new(user)

type user struct{}

func (u *user) Login(c *gin.Context) {
	var req types.LoginReq

	// TODO: req绑定和验证，调用service方法

	c.JSON(http.StatusOK, req)
	return
}
