package ini

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/ini.v1"
)

func Main() (string, int) {
	//[ini 配置]
	cfg, err := ini.Load("./conf/server.ini")
	if err != nil{
		logs.Critical("No conf/server.ini", err.Error())
		return "", 0
	}
	sec, err := cfg.GetSection("web")
	if err != nil{
		logs.Critical("No web section in conf/server.ini", err.Error())
		return "", 0
	}
	default_url, err := sec.GetKey("default_url")
	url := "/"
	if err == nil {
		url = default_url.String()
	}
	port, err := sec.GetKey("port")
	if err != nil {
		logs.Critical("No port value in conf/server.ini", err.Error())
		return "", 0
	}
	webport, err := port.Int()
	if err != nil{
		logs.Critical("port not number", err.Error())
		return "", 0
	}

	return url, webport

}
