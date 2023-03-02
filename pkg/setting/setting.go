package setting

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Server struct{
	RunMode string
	HttpPort int
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}



var cfg *ini.File

// 读取conf配置文件，并赋值变量
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("fail to read file: %v", err)
	}

	cfg.Section("server").MapTo(ServerSetting)
	cfg.Section("database").MapTo(DatabaseSetting)

}