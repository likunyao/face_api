package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

type Application struct {
	LogSavePath string
	LogFileExt  string
	TimeFormat  string
	JwtSecret   string
}

var DatabaseSetting = &Database{}
var ServerSetting = &Server{}
var ApplicationSetting = &Application{}

func Setup() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config.ini':%v", err)
	}
	mapTo("database", DatabaseSetting, cfg)
	mapTo("server", ServerSetting, cfg)
	mapTo("application", ApplicationSetting, cfg)
}

func mapTo(section string, v interface{}, cfg *ini.File) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.Mapto %s err: %v", section, err)
	}
}
