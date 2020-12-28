package controllers

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/yinrenxin/hgoj/syserror"
	"github.com/beego/beego/v2/adapter/logs"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if c.IsAjax() {
		c.Ctx.Output.Status = 200
		c.Data["json"] = map[string]interface{}{
			"code": 1002,
			"msg": "非法访问",
		}
		c.ServeJSON()
	}
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	err, ok := c.Data["error"].(error)

	if !ok {
		err = syserror.New("未知错误", nil)
	}

	serr, ok := err.(syserror.Error)
	if !ok {
		serr = syserror.New(err.Error(), nil)
	}

	if serr.ReasonError() != nil {
		logs.Info(serr.Error(),serr.ReasonError())
	}

	if c.IsAjax() {
		c.jsonerror(serr)
	} else {
		c.Data["content"] = serr.Error()
	}
}

func (c *ErrorController) Error401() {
	c.TplName = "error/401.html"
	err, ok := c.Data["error"].(error)

	if !ok {
		err = syserror.New("未知错误", nil)
	}

	serr, ok := err.(syserror.Error)
	if !ok {
		serr = syserror.New(err.Error(), nil)
	}

	if serr.ReasonError() != nil {
		logs.Info(serr.Error(),serr.ReasonError())
	}

	if c.IsAjax() {
		c.jsonerror(serr)
	} else {
		c.Data["content"] = serr.Error()
	}
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}

func (c *ErrorController) jsonerror(serr syserror.Error){
	c.Ctx.Output.Status = 200
	c.Data["json"] = map[string]interface{}{
		"code":serr.Code(),
		"msg":serr.Error(),
	}
	c.ServeJSON()
}
