package controllers

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/yinrenxin/hgoj/models"

	"github.com/beego/beego/v2/adapter/logs"
	types "github.com/yinrenxin/hgoj/models/types"
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
	k := time.Now().Unix()
	dir_name := OJ_DATA + "/exporttemp/" + strconv.Itoa(int(k))
	filename := dir_name + "/" + "export.json"
	os.MkdirAll(dir_name, os.ModePerm)
	var pros []types.Problem
	for i := from; i <= to; i++ {
		modelpro, err := models.QueryProblemById(int32(i))
		if err != nil {
			continue
		}

		pro := types.Problem{
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
		cmd := exec.Command("cp", "-r", pro_dir, dir_name)
		err2 := cmd.Run()
		if err2 != nil {
			logs.Error("Execute Command failed:" + err2.Error())
			return
		}
	}
	json_data, _ := json.Marshal(pros)
	err1 := ioutil.WriteFile(filename, json_data, os.ModePerm)
	if err1 != nil {
		logs.Error(err1)
	}

	exportDataDir, _ := os.Open(dir_name)
	zipdestDir := DOWN_DIR + "export-" + strconv.Itoa(from) + "-" + strconv.Itoa(to) + "-" + strconv.Itoa(int(k)) + ".zip"

	files := []*os.File{exportDataDir}
	err := tools.Compress(files, zipdestDir)
	if err != nil {
		logs.Error(err)
	}
	downurl := "/static/down/" + "export-" + strconv.Itoa(from) + "-" + strconv.Itoa(to) + "-" + strconv.Itoa(int(k)) + ".zip"
	resData := MAP_H{
		"data": MAP_H{
			"downurl": downurl,
		},
	}
	this.JsonOKH("导出成功", resData)
}

// @router /problem/inport [post]
func (this *ProblemController) Inport() {
	if !this.IsAdmin {
		this.Abort("401")
	}

	//获取上传的数据
	f, h, err := this.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	defer f.Close()

	//临时保存数据
	this.SaveToFile("file", OJ_ZIP_TEMP_DATA+"/"+h.Filename)
	uploadZipData := OJ_ZIP_TEMP_DATA + "/" + h.Filename
	err = tools.DeCompress(uploadZipData, OJ_ZIP_TEMP_DATA)
	if err != nil {
		logs.Error(err)
		this.JsonErr(err.Error(), 24001, "")
	}
	err = os.Remove(uploadZipData)
	if err != nil {
		logs.Error(err)
		this.JsonErr(err.Error(), 24002, "")
	}

	//获取解压后的目录
	fileHashName := (h.Filename)[len(h.Filename)-14 : len(h.Filename)-4]
	fileDir := OJ_ZIP_TEMP_DATA + "/" + fileHashName
	jsonDataDir := fileDir + "/export.json"

	//读取json数据到struct
	fileData, err := ioutil.ReadFile(jsonDataDir)
	if err != nil {
		logs.Error(err)
		this.JsonErr(err.Error(), 24002, "")
	}
	var problems []types.Problem
	err = json.Unmarshal(fileData, &problems)
	if err != nil {
		logs.Error(err)
		this.JsonErr(err.Error(), 24002, "")
	}

	//处理读出来的数据(如果title一样跳过)
	var repeatProblems []types.Problem
	var insertProblems []types.Problem
	for _, v := range problems {
		_, err := models.QueryProblemByTitle(v.Title)
		if err == nil {
			repeatProblems = append(repeatProblems, v)
			continue
		}
		insertProblems = append(insertProblems, v)
	}

	//插入数据
	pidMap, err := models.BatchAddProblem(insertProblems)
	if err != nil {
		logs.Error(err)
		this.JsonErr(err.Error(), 24002, "")
	}

	//移动题目数据
	for k, v := range pidMap {
		testDataDir := OJ_DATA + "/" + strconv.Itoa(int(v))
		oldtestDataDir := fileDir + "/" + strconv.Itoa(int(k))
		logs.Info(testDataDir)
		logs.Info(oldtestDataDir)
		os.Rename(oldtestDataDir, testDataDir)

	}

	var d []MAP_H
	for _, v := range insertProblems {
		d = append(d, MAP_H{
			"pid":   pidMap[v.ProblemId],
			"title": v.Title,
		})
	}

	resData := MAP_H{
		"data": d,
	}

	this.JsonOKH("上传成功", resData)
}
