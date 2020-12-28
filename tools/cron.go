package tools

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/toolbox"
	"io/ioutil"
	"os"
	"time"
)

func StartCron() {
	tk := toolbox.NewTask("clearDownData", "0 0 */1 * * *", func() error { clearDownData(); return nil })
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("clearDownData", tk)
	toolbox.StartTask()
	//time.Sleep(6 * time.Second)
	//toolbox.StopTask()
}


func clearDownData() {
	downDir := "./static/down"
	zipDir := "/home/judge/data/tempzip"

	dir_list, e := ioutil.ReadDir(downDir)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for _, v := range dir_list {
		finfo, _ := os.Stat(downDir+"/"+v.Name())
		var fctime time.Time
		fctime = finfo.ModTime()
		t := time.Now().Sub(fctime).Minutes()
		logs.Info(t)
		if t > 10 {
			err := os.Remove(downDir+"/"+v.Name())
			if err != nil {
				logs.Error(err)
			}
		}
		logs.Info("delete file ",downDir+"/"+v.Name())
	}


	zipdir_list, e1 := ioutil.ReadDir(zipDir)
	if e1 != nil {
		logs.Error("read dir error")
		return
	}
	for _, v1 := range zipdir_list {
		finfo, _ := os.Stat(zipDir+"/"+v1.Name())
		var fctime time.Time
		fctime = finfo.ModTime()
		t := time.Now().Sub(fctime).Minutes()
		logs.Info(t)
		if t > 10 {
			err := os.Remove(zipDir+"/"+v1.Name())
			if err != nil {
				logs.Error(err)
			}
		}
		logs.Info("delete file ",zipDir+"/"+v1.Name())
	}
}
