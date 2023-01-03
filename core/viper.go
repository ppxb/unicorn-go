package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitViper() {
	var config string
	var mode string

	switch gin.Mode() {
	case gin.DebugMode:
		config = ConfigDefaultFile
		mode = gin.DebugMode
	case gin.TestMode:
		config = ConfigTestFile
		mode = gin.TestMode
	case gin.ReleaseMode:
		config = ConfigReleaseFile
		mode = gin.ReleaseMode

	}
	fmt.Printf("项目环境为：%s，配置文件为：%s \n", mode, config)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败：%s", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被改变了：", in.Name)
		injectConfig(v)
	})

	injectConfig(v)
}

func injectConfig(v *viper.Viper) {
	if err := v.Unmarshal(&AppConfig); err != nil {
		fmt.Println(err)
	}
}
