package controllers

import (
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/yinrenxin/hgoj/syserror"

	//"github.com/beego/beego/v2/adapter/logs"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
)

type ProblemController struct {
	BaseController
}

type Problems struct {
	num       int64
	pageRange []int
	pageNum   int
}

// @router /problemset/:page [get]
func (this *ProblemController) ProblemSetPage() {
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo) * pageSize
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
	pagePrev := pageNo
	pageNext := pageNo + 2
	if int(pageNo) == temp {
		pageNext = pageNo + 1
	}
	if pageNo == 0 {
		pagePrev = pageNo + 1
	}

	this.Data["pageData"] = proData
	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["problems"] = pros
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.TplName = "problem.html"
}

// @router /problem/:id [get]
func (this *ProblemController) Problem() {
	id := this.Ctx.Input.Param(":id")
	ids, err := tools.StringToInt32(id)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}
	pro, err := models.QueryProblemById(ids)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}

	if !this.IsAdmin && pro.Defunct == "Y" {
		this.Abort("401")
	}

	this.Data["problem"] = pro
	this.TplName = "problems.html"
}

// @router /problem/edit/:id [get]
func (this *ProblemController) ProblemEdit() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	id := this.Ctx.Input.Param(":id")
	ids, err := tools.StringToInt32(id)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}
	pro, err := models.QueryProblemById(ids)
	if err != nil {
		this.Abort("401")
		logs.Error(err)
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["PRO"] = pro
	this.TplName = "admin/editProblem.html"
}

// @router /problem/update [post]
func (this *ProblemController) ProblemUpdate() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	proId := this.GetString("proid")
	temp, err := strconv.Atoi(proId)
	id := int32(temp)

	pro, err := models.QueryProblemById(id)
	if err != nil {
		this.JsonErr("问题未找到", syserror.PROBLEM_NOT_FOUND, "/problem/edit/"+proId)
		logs.Error(pro, err)
	}

	title := this.GetMushString("title", "标题不能为空")
	memory := this.GetMushString("memory", "限制内存不能为空")
	rtime := this.GetMushString("time", "限制时间不能为空")
	desc := this.GetMushString("desc", "描述不能为空")
	input := this.GetMushString("input", "input不能为空")
	output := this.GetMushString("output", "output不能为空")
	sampleinput := this.GetMushString("sampleinput", "sampleinput不能为空")
	sampleoutput := this.GetMushString("sampleoutput", "sampleoutput不能为空")
	spj := this.GetMushString("spj", "spj不能为空")
	hint := this.GetString("hint")
	inDate := time.Now()
	data := []string{title, rtime, memory, desc, input, output, sampleinput, sampleoutput, hint, spj}
	ok, err := models.UpdateProblemById(id, data, inDate)
	if !ok {
		this.JsonErr("更新失败", syserror.UPDATE_PROBLEM_ERR, "problem/edit/"+proId)
	}

	this.JsonOK("更新成功", "/problem/list")
}

// @router /problem/updatestatus [post]
func (this *ProblemController) ProblemUpdateStatus() {
	if !this.IsAdmin {
		this.Abort("401")
	}

	temp := this.GetString("proid")
	pid, _ := tools.StringToInt32(temp)
	if ok := models.UpdateProStatus(pid); !ok {
		this.JsonErr("失败", 16001, "")
	}

	this.JsonOK("成功", "")
}

// @router /problem/add [get]
func (this *ProblemController) ProblemAdd() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/addProblem.html"
}

// @router /problem/del [post]
func (this *ProblemController) ProblemDel() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	this.JsonErr("暂不开放删除功能", syserror.DEL_PROBLEM_ERR, "/problem/list")
	proId := this.GetString("proid")
	temp, _ := strconv.Atoi(proId)
	id := int32(temp)
	ok := models.DelProblemById(id)
	if !ok {
		this.JsonErr("删除失败", syserror.DEL_PROBLEM_ERR, "/problem/list")
	}
	err2 := os.RemoveAll(OJ_DATA + "/" + proId)
	if err2 != nil {
		this.JsonErr("删除失败", syserror.DEL_PROBLEM_ERR, "/problem/list")
	}
	this.JsonOK("删除成功", "/problem/list")
}

// @router /problem/list [get]
func (this *ProblemController) ProblemList() {
	if !this.IsAdmin && !this.IsTeacher {
		this.Abort("401")
	}
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
	pageRange := t
	pagePrev := pageNo + 1
	pageNext := pageNo + 2

	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["PROS"] = pros
	this.TplName = "admin/listProblem.html"
}

// @router /problem/list/:page [get]
func (this *ProblemController) ProblemListPage() {
	if !this.IsAdmin && !this.IsTeacher {
		this.Abort("401")
	}
	page := this.Ctx.Input.Param(":page")
	pageNo, _ := tools.StringToInt32(page)
	pageNo = pageNo - 1
	start := int(pageNo) * pageSize
	pros, _, totalNum, err := models.QueryPageProblem(start, pageSize)
	if err != nil {
		logs.Error(err)
	}

	isPage, pageRange, pagePrev, pageNext := PageRangeCal(totalNum, pageNo, pageSize)

	this.Data["pageRange"] = pageRange
	this.Data["isPage"] = isPage
	this.Data["pagePrev"] = pagePrev
	this.Data["pageNext"] = pageNext
	this.Data["PROS"] = pros
	this.TplName = "admin/listProblem.html"
}

// @router /problem/add [post]
func (this *ProblemController) ProblemAddPost() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	title := this.GetMushString("title", "标题不能为空")
	memory := this.GetMushString("memory", "限制内存不能为空")
	rtime := this.GetMushString("time", "限制时间不能为空")
	desc := this.GetMushString("desc", "描述不能为空")
	input := this.GetMushString("input", "input不能为空")
	output := this.GetMushString("output", "output不能为空")
	sampleinput := this.GetMushString("sampleinput", "sampleinput不能为空")
	sampleoutput := this.GetMushString("sampleoutput", "sampleoutput不能为空")
	spj := this.GetMushString("spj", "spj不能为空")
	hint := this.GetString("hint")
	filename := this.GetMushString("filename", "未上传测试数据")
	inDate := time.Now()
	data := []string{title, rtime, memory, desc, input, output, sampleinput, sampleoutput,hint,spj}
	pid, err := models.AddProblem(data, inDate)
	if err != nil {
		this.JsonErr("更新失败", syserror.ADD_PROBLEM_ERR, "/problem/add")
	}
	testDataDir := OJ_DATA + "/" + strconv.Itoa(int(pid)) + "/"
	zipDataDir := OJ_ZIP_TEMP_DATA + "/" + filename
	err1 := tools.DeCompress(zipDataDir, testDataDir)
	if err1 != nil {
		logs.Error(err1)
		if ok := models.DelProblemById(int32(pid)); ok {
			logs.Error("delete err problem")
		}
		this.JsonErr("系统错误", 2300, "")
	}
	err2 := os.Remove(zipDataDir)
	if err2 != nil {
		logs.Error(err2)
		if ok := models.DelProblemById(int32(pid)); ok {
			logs.Error("delete err problem")
		}
		this.JsonErr("系统错误", 2300, "")
	}
	this.JsonOK("添加题目成功", "/problem/list")
}

// @router /problem/fileupload [post]
func (this *ProblemController) Fileupload() {
	if !this.IsAdmin && !this.IsTeacher {
		this.Abort("401")
	}
	key := tools.MD5(time.Now().String())
	f, h, err := this.GetFile("file")
	//if h.Filename != "data.zip" {
	//	this.JsonErr("文件名错误,请上传zip压缩包并命名为data.zip",2400,"")
	//}

	if err != nil {
		logs.Error("error:--- ", err)
	}
	defer f.Close()

	zipDir := OJ_ZIP_TEMP_DATA
	err2 := os.Mkdir(zipDir, os.ModePerm)
	if err2 != nil {
		//logs.Error(err2)
	}

	err1 := this.SaveToFile("file", OJ_ZIP_TEMP_DATA+"/"+key+h.Filename)
	if err1 != nil {
		this.JsonErr("文件上传错误", 24005, "")
	}
	data := MAP_H{
		"key":      key,
		"filename": h.Filename,
	}
	this.JsonOKH("上传成功", data)

}

func mkdata(pid int64, filename string, input string, oj_data string) bool {
	baseDir := oj_data + "/" + strconv.Itoa(int(pid))
	err := os.MkdirAll(baseDir, 0777)
	if err != nil {
		logs.Error("目录创建失败", err)
		return false
	}
	name := baseDir + "/" + filename
	data := []byte(input)
	if ioutil.WriteFile(name, data, 0644) != nil {
		logs.Error("文件写入失败,文件名", name)
		return false
	}
	return true
}
