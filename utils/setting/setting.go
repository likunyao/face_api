package setting

import (
	"github.com/go-ini/ini"
	"log"
)

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	cfg, err := ini.Load("../config/config.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/config.ini':%v",err)
	}
	mapTo("database", DatabaseSetting, cfg)
}

func mapTo(section string, v interface{}, cfg *ini.File) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.Mapto %s err: %v", section, err)
	}
}