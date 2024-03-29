package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ini/ini"
	"github.com/spf13/viper"
)

var (
	Cfg *ini.File

	HTTPPort int
	RunMode  string

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

// 获取配置文件
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini':%v", err)
	}

	LoadBase()
	LoadServer()
	LoadServer()
	viper.SetConfigName("config")
	viper.AddConfigPath("conf/")

	err2 := viper.ReadInConfig()
	if err2 != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	str := viper.Get("unbound_server").(map[string]interface{})
	fmt.Println("viper", str["port"])
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(3000)

	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	//拿到[app]配置
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
