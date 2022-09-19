package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"github.com/yinrenxin/hgoj/models"
	"github.com/yinrenxin/hgoj/tools"
	"testing"
)

func TestSaveUser(t *testing.T) {
	initTemplate()
	initSession()
	initStatic()
	initLogs()
	//启动定时任务
	go tools.StartCron()
	go tools.InitTools()

	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		username := row[5]
		nick := row[4] + row[2]
		pwd := row[5]
		_, err := models.SaveUser(username, nick, "", pwd, "湖南工学院", "")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(nick, pwd)
	}

}
