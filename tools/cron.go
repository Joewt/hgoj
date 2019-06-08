package tools

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

func StartCron() {

	t1 := time.NewTimer(time.Second * 3600)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 3600)
			clearDownData()
		}
	}
}


func clearDownData() {
	logs.Info("执行了")
	downDir := "./static/down"
	dir_list, e := ioutil.ReadDir(downDir)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	for _, v := range dir_list {
		finfo, _ := os.Stat(downDir+"/"+v.Name())
		stat_t := finfo.Sys().(*syscall.Stat_t)
		fctime := timespecToTime(stat_t.Ctimespec)
		t := time.Now().Sub(fctime).Minutes()
		if t > 10 {
			err := os.Remove(downDir+"/"+v.Name())
			if err != nil {
				logs.Error(err)
			}
		}
		logs.Info("delete file ",downDir+"/"+v.Name())
	}
}


func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}