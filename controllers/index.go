package controllers

import (
	"encoding/json"
	"html/template"
	"time"

	"github.com/go-redis/redis"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/syserror"
)

var store = cache.NewMemoryCache()
var CPT = captcha.NewWithFilter("/captcha/", store)

func init() {
	CPT.ChallengeNums = 4
	CPT.StdWidth = 120
	CPT.StdHeight = 40
}

type IndexController struct {
	BaseController
}

var (
	pageSize int = 50
	//pageProSize int = 100
	pageStatusSize  int = 40
	pageContestSize int = 20
)

// @router / [get]
func (this *IndexController) Indexs() {
	this.TplName = "empty.html"
}

// @router /index [get]
func (this *IndexController) Index() {
	_, ok := this.GetSession(SESSION_USER_KEY).(models.Users)
	if !ok {
		logs.Error("未登陆")
	}
	//user, _, err := models.QueryAllUser()
	//if err != nil {
	//	this.JsonErr("未知错误", 4000, "/index")
	//}

	//cache.NewCache("memory", `{"key":"hgoj","conn":"r-8vb4cmly4tyog47ykepd.redis.zhangbei.rds.aliyuncs.com:6379","dbNum":"2","":""}`)
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//sort.Sort(SortUser(user))

	//var RankUser []*RankUsers
	//
	//for k, v := range user {
	//	RankUser = append(RankUser, &RankUsers{k + 1, *v})
	//}

	//if len(RankUser) > 30 {
	//	RankUser = RankUser[0:20]
	//}

	art, err := models.QueryLimitArt()
	if err != nil {
		logs.Error("没有文章")
	}
	nowTime := time.Now().Format("2006-01-02")
	totalNum, acNum := models.QueryTotalNumAcNumSolution(nowTime)

	var totalNs []int64
	var acNs []int64
	var times []string
	type TAC struct {
		totalNs []int64
		acNs    []int64
		times   []string
	}
	var redisKeys = "index" + nowTime
	logs.Error(redisKeys)
	val, err := client.Get(redisKeys).Result()
	if err != nil {
		logs.Error("redic get errr:", err)
	}
	if val == "" {
		for i := -7; i <= -1; i++ {
			calTime := time.Now().AddDate(0, 0, i).Format("2006-01-02")
			totalN, acN := models.QueryTotalNumAcNumSolution(calTime)
			totalNs = append(totalNs, totalN)
			acNs = append(acNs, acN)
			times = append(times, calTime)
		}
		var tac *TAC
		tac = &TAC{totalNs: totalNs, acNs: acNs, times: times}
		json_data, _ := json.Marshal(tac)
		_ = client.SetNX(redisKeys, json_data, 3600*time.Second).Err()
	}
	res, err := client.Get(redisKeys).Result()
	var redisTotalNsA []int64
	var redisAcNs []int64
	var redisTimes []string
	if err == nil {
		var tacs *TAC
		err = json.Unmarshal([]byte(res), &tacs)
		if err == nil {
			logs.Error("tacs", tacs)
			redisTotalNsA = tacs.totalNs
			redisAcNs = tacs.acNs
			redisTimes = tacs.times
		}
	}


	//this.Data["user"] = RankUser
	this.Data["totalNum"] = totalNum
	this.Data["acNum"] = acNum
	this.Data["nowTime"] = nowTime
	this.Data["Art"] = art
	this.Data["totalNums"] = redisTotalNsA
	this.Data["acNums"] = redisAcNs
	this.Data["times"] = redisTimes
	this.TplName = "index.html"
}

// @router /faqs [get]
func (this *IndexController) IndexFaqs() {
	this.TplName = "faqs.html"
}

// @router /problemset [get]
func (this *IndexController) IndexProblemset() {

	//_,num,err := models.QueryAllProblem()
	//if err != nil {
	//	logs.Error(err)
	//}
	pageNo := 0
	start := pageNo * pageSize
	pros, _, totalNum, err := models.QueryPageProblem(start, pageSize)
	if err != nil {
		logs.Error(err)
	}

	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	temp := int(totalNum) / pageSize
	var t []int
	for i := 0; i <= temp; i++ {
		t = append(t, i+1)
	}
	proData := new(Problems)
	proData.pageRange = t
	proData.num = totalNum
	pageRange := t

	pagePrev := pageNo + 1
	pageNext := pageNo + 2

	this.Data["pageData"] = proData
	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["problems"] = pros
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.TplName = "problem.html"
}

// @router /status [get]
func (this *IndexController) IndexStatus() {
	pageNo := 0
	start := int(pageNo) * pageSize
	data, RESULT, _, totalNum, err := models.QueryPageSolution(start, pageSize)
	if err != nil {
		logs.Error(err)
	}
	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	pagePrev := 1
	pageNext := 2
	this.Data["data"] = data
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["RES"] = RESULT
	this.TplName = "status.html"
}

// @router /contest [get]
func (this *IndexController) IndexContest() {
	pageNo := 0
	start := int(pageNo) * pageContestSize
	con, _, totalNum, _ := models.QueryPageContest(start, pageContestSize)
	isPage := true
	if int(totalNum) < pageContestSize {
		isPage = false
	}
	pagePrev := 1
	pageNext := 2
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["con"] = con
	this.TplName = "contest.html"
}

// @router /login [get]
func (this *IndexController) IndexUser() {
	if this.IsLogin {
		this.Abort("401")
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["captcha"] = true
	this.TplName = "login.html"
}

// @router /reg [get]
func (this *IndexController) IndexReg() {
	if this.IsLogin {
		this.Redirect("/index", 302)
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["captcha"] = true
	this.TplName = "reg.html"
}

// @router /admin [get]
func (this *IndexController) IndexAdmin() {
	if !this.IsAdmin && !this.IsTeacher {
		this.Abort401(syserror.UnKnowError{})
	}
	//v, _ := mem.VirtualMemory()
	//this.Data["memused"] = (v.Total / (1024 * 1024)) - (v.Free / (1024 * 1024))
	//this.Data["memfree"] = v.Free / (1024 * 1024)
	this.Data["memused"] = 0
	this.Data["memfree"] = 0
	this.TplName = "admin/index.html"
}

//
//// @router /skin-config.html [get]
//func (this *IndexController) IndexSkinConfig() {
//	this.TplName = "admin/skin-conf.html"
//}
