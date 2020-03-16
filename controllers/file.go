package controllers

import (
	"encoding/json"
	"github.com/yinrenxin/hgoj/models"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/tools"
)

const DOWN_DIR = "./static/down/"

// @router /problem/editdata/:pid [get]
func (this *ProblemController) ProblemTestDataEdit() {

	temp := this.Ctx.Input.Param(":pid")

	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["pid"] = temp
	this.TplName = "admin/editData.html"
}

// @router /download/testdata [post]
func (this *ProblemController) DownloadTestData() {
	if !this.IsAdmin {
		this.Abort("401")
	}

	pid := this.GetMushString("pid", "系统错误")
	id := tools.StringToInt(pid)
	testDataDir, _ := os.Open(OJ_DATA + "/" + strconv.Itoa(id) + "/")
	zipdestDir := DOWN_DIR + "data-" + pid + ".zip"

	files := []*os.File{testDataDir}
	err := tools.Compress(files, zipdestDir)
	if err != nil {
		logs.Error(err)
	}
	downurl := "/static/down/" + "data-" + pid + ".zip"
	this.JsonOK("下载成功", downurl)
}

// @router /upload/fileupload/:pid [post]
func (this *ProblemController) Upload() {
	pid := this.Ctx.Input.Param(":pid")
	if !this.IsAdmin {
		this.Abort("401")
	}
	id, _ := tools.StringToInt32(pid)
	key := tools.MD5(time.Now().String())
	f, h, err := this.GetFile("file")
	//if h.Filename != "data.zip" {
	//	this.JsonErr("文件名错误,请上传zip压缩包并命名为data.zip",2400,"")
	//}

	if err != nil {
		logs.Error("error:--- ", err)
	}
	defer f.Close()
	this.SaveToFile("file", OJ_ZIP_TEMP_DATA+"/"+key+h.Filename)

	testDataDir := OJ_DATA + "/" + strconv.Itoa(int(id)) + "/"
	zipDataDir := OJ_ZIP_TEMP_DATA + "/" + key + h.Filename

	err2 := os.RemoveAll(testDataDir)

	if err2 != nil {
		logs.Error(err)
		this.JsonErr("系统错误，请查看系统日志", 24002, "")
	}

	err1 := tools.DeCompress(zipDataDir, testDataDir)
	if err1 != nil {
		logs.Error(err1)
		this.JsonErr("系统错误，请查看系统日志", 24001, "")
	}

	err3 := os.Remove(zipDataDir)
	if err3 != nil {
		logs.Error(err3)
		this.JsonErr("系统错误，请查看系统日志", 24002, "")
	}
	data := MAP_H{
		"key":      key,
		"filename": h.Filename,
	}
	this.JsonOKH("上传成功", data)

}

// @router /problem/exinport [get]
func (this *ProblemController) ExInport() {

	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "admin/exinport.html"
}

// @router /problem/export [post]
func (this *ProblemController) Export() {
	if !this.IsAdmin {
		this.Abort("401")
	}
	from, _ := this.GetInt("from")
	to, _ := this.GetInt("to")
	if from > to {
		this.JsonErr("导出错误,from小于to", 3100, "")
	}
	type Problem struct {
		ProblemId		int32		`json:"problem_id"`
		Title			string		`json:"title"`
		Description 	string		`json:"description"`
		Input			string		`json:"input"`
		Output			string		`json:"output"`
		SampleInput		string		`json:"sampleinput"`
		SampleOutput	string		`json:"sampleoutput"`
		Spj				string		`json:"spj"`
		Hint			string		`json:"hint"`
		Source			string		`json:"source"`
		InDate			time.Time	`json:"in_date"`
		TimeLimit		int32		`json:"time_limit"`
		MemoryLimit		int32		`json:"memory_limit"`
		Defunct			string		`json:"defunct"`
		Accepted		int32		`json:"accepted"`
		Submit			int32		`json:"submit"`
		Solved			int32		`json:"solved"`
	}
	//
	k := time.Now().Unix()
	dir_name := OJ_DATA+"/exporttemp/"+strconv.Itoa(int(k))
	filename := dir_name+"/"+"export.json"
	os.MkdirAll(dir_name,os.ModePerm)
	var pros []Problem
	for i := from; i <= to; i++ {
		modelpro,err := models.QueryProblemById(int32(i))
		if err != nil {
			continue
		}

		pro := Problem{
			ProblemId:    modelpro.ProblemId,
			Title:        modelpro.Title,
			Description:  modelpro.Description,
			Input:        modelpro.Input,
			Output:       modelpro.Output,
			SampleInput:  modelpro.SampleInput,
			SampleOutput: modelpro.SampleOutput,
			Spj:          modelpro.Spj,
			Hint:         modelpro.Hint,
			Source:       modelpro.Source,
			InDate:       modelpro.InDate,
			TimeLimit:    modelpro.TimeLimit,
			MemoryLimit:  modelpro.MemoryLimit,
			Defunct:      modelpro.Defunct,
			Accepted:     modelpro.Accepted,
			Submit:       modelpro.Submit,
			Solved:       modelpro.Solved,
		}
		pros = append(pros, pro)
		pro_dir := OJ_DATA + "/" + strconv.Itoa(int(modelpro.ProblemId))
		cmd := exec.Command("cp", "-r",pro_dir , dir_name)
		err2 := cmd.Run()
		if err2 != nil {
			logs.Error("Execute Command failed:" + err2.Error())
			return
		}
	}
	json_data, _ := json.Marshal(pros)
	err1 := ioutil.WriteFile(filename,json_data,os.ModePerm)
	if err1 != nil {
		logs.Error(err1)
	}

	exportDataDir, _ := os.Open(dir_name)
	zipdestDir := DOWN_DIR + "export-" + strconv.Itoa(from)+"-"+strconv.Itoa(to)+ ".zip"

	files := []*os.File{exportDataDir}
	err := tools.Compress(files, zipdestDir)
	if err != nil {
		logs.Error(err)
	}
	downurl := "/static/down/" + "export-" + strconv.Itoa(from)+"-"+strconv.Itoa(to)+ ".zip"
	 resData := MAP_H{
		"data": MAP_H{
			"downurl": downurl,
		},
	}
	this.JsonOKH("导出成功", resData)
}

// @router /problem/inport [post]
func (this *ProblemController) Inport() {

	this.JsonErr("未开放", 2333, "")
	//var prob models.Problem
	//json.Unmarshal([]byte(jsonPro),&prob)
	//logs.Info(i,"-----",prob)
}
