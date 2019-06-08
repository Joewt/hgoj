package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/yinrenxin/hgoj/tools"
	"os"
	"strconv"
)

const DOWN_DIR = "./static/down/"

// @router /problem/editdata/:pid [get]
func (this *ProblemController) ProblemTestDataEdit() {


	temp := this.Ctx.Input.Param(":pid")

	this.Data["pid"] = temp
	this.TplName = "admin/editData.html"
}


// @router /download/testdata [post]
func (this *ProblemController) DownloadTestData(){
	if !this.IsAdmin {
		this.Abort("401")
	}

	pid := this.GetMushString("pid","系统错误")
	id := tools.StringToInt(pid)
	testDataDir,_ := os.Open(OJ_DATA+"/"+strconv.Itoa(id)+"/")
	zipdestDir := DOWN_DIR + "data-"+pid+".zip"

	files := []*os.File{testDataDir}
	err := tools.Compress(files,zipdestDir)
	if err != nil {
		logs.Error(err)
	}
	downurl := "/static/down/"+"data-"+pid+".zip"
	this.JsonOK("下载成功",downurl)
}


// @router /upload/fileupload/:pid [post]
func (this *ProblemController) Upload() {
	pid := this.Ctx.Input.Param(":pid")
	if !this.IsAdmin{
		this.Abort("401")
	}
	id,_ := tools.StringToInt32(pid)
	//key := tools.MD5(time.Now().String())
	f, h, err := this.GetFile("file")
	if h.Filename != "data.zip" {
		this.JsonErr("文件名错误,请上传zip压缩包并命名为data.zip",2400,"")
	}

	if err != nil {
		logs.Error("error:--- ",err)
	}
	defer f.Close()
	this.SaveToFile("file", OJ_DATA +"/"+h.Filename)


	testDataDir := OJ_DATA+"/"+strconv.Itoa(int(id))+"/"
	zipDataDir := OJ_DATA+"/"+h.Filename

	err2 := os.RemoveAll(testDataDir)

	if err2 != nil {
		logs.Error(err)
		this.JsonErr("err",24002,"")
	}


	err1 := tools.DeCompress(zipDataDir, testDataDir)
	if err1 != nil {
		logs.Error(err1)
		this.JsonErr("系统错误",24001,"")
	}

	err3 := os.Remove(zipDataDir)
	if err3 != nil {
		logs.Error(err3)
		this.JsonErr("系统错误",24002,"")
	}
	data := MAP_H{
		//"key":key,
		"filename":h.Filename,
	}
	this.JsonOKH("上传成功",data)

}