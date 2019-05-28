package test

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"github.com/yinrenxin/hgoj/controllers"
)

func prepareController(c *beego.Controller) {
	c.Ctx = &context.Context{
		Request:        &http.Request{URL: &url.URL{Scheme: "http", Host: "localhost:8086", Path: "/"}},
		ResponseWriter: &fakeResponseWriter{},
	}
	c.Ctx.Output = &context.BeegoOutput{Context: c.Ctx}
	c.Ctx.Input = &context.BeegoInput{Request: c.Ctx.Request}

	globalSessions, _ := session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":10}`)
	c.Ctx.Request.Header = http.Header{}
	c.Ctx.Request.AddCookie(&http.Cookie{Name: "gosessionid", Value: "test"})
	c.CruSession = globalSessions.SessionRegenerateId(c.Ctx.ResponseWriter, c.Ctx.Request)
	c.Data = map[interface{}]interface{}{}
}

func TestRecomputeBanlance(t *testing.T) {
	c := &controllers.BanlanceController{}
	prepareController(&(c.Controller))
	// 这是期望用户在浏览器上传的form表和登陆信息。
	c.Ctx.Request.Form = url.Values{
		"range": []string{"2016-10-01到2016-10-31"},
		"city": []string{"北京"},
	}
	c.SetSession("login", "123")
	c.Prepare()
	// 这是对应的控制器函数
	c.Recompute()
	// 本例中，数据是通过json传出来的。
	j := c.Data["json"]
	// 从json中读取的数据是interface{}类型，需要通过reflect获取其中的信息
	mapV := reflect.ValueOf(j)
	errV := mapV.MapIndex(reflect.ValueOf("Error"))
	// 先看看json中有没有error字段
	if errV.IsValid() && errV.String() != "" {
		t.Fatal("has error:", errV)
	}
	// 然后读取ban字段的内容
	banV := mapV.MapIndex(reflect.ValueOf("ban"))
	if !banV.IsValid() {
		t.Fatal("no output data!")
	} else {
		V := reflect.ValueOf(banV.Interface())
		fnum := V.NumField()
		if fnum < 10 {
			t.Fatal("not ban table format")
		}
		// 读取关键字段的值进一步判断
		v1 := V.FieldByName("Value1").Float()
		v2 := V.FieldByName("Value2").Float()
		v3 := V.FieldByName("Value3").Float()
		v4 := V.FieldByName("Value4").Float()
		if !(v1 > 0 &&
			v2 > 0 &&
			v3 > 0 &&
			v4 > 0) {
			t.Fatal("ban table data wrong:", v1, v2, v3, v4)
		}
	}
}

type fakeResponseWriter struct{}

func (f *fakeResponseWriter) Header() http.Header {
	return http.Header{}
}
func (f *fakeResponseWriter) Write(b []byte) (int, error) {
	return 0, nil
}
func (f *fakeResponseWriter) WriteHeader(n int) {}