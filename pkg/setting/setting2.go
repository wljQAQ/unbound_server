package setting

import "github.com/spf13/viper"

var (
	Server ServerConf
)

type ServerConf struct {
	Port string `mapstructure:"port"`
	Db   struct {
		Type     string `mapstructure:"type"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"db"`
}

func init() {
	// 设置配置文件的名称
	viper.SetConfigName("config")
	// 设置配置文件的路径
	viper.AddConfigPath("conf/")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 解析配置文件中的 "unbound_server" 键，并将其值赋给 Server 变量
	if err := viper.UnmarshalKey("unbound_server", &Server); err != nil {
		panic(err)
	}
}
