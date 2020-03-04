package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/sunnu06/personal-website/models"
	"github.com/sunnu06/personal-website/routers"
)

func main()  {
	//[logs 配置]
	logs.SetLogger(logs.AdapterConsole) // "console"
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"log/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"separate":["emergency"]}`)
	logs.Async(1000)
	logs.Info("set logger ok")

	//[mysql 配置]
	models.Init()
	//[Router 初始值配置]
	routers.Main()

}