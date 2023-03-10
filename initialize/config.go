package initialize

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/ppxb/unicorn/models"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/log"
	"github.com/spf13/viper"
	"strings"
)

const (
	configType            = "yaml"
	configDir             = "conf"
	configFile            = "config.yml"
	defaultConnectTimeout = 5
	defaultMaxIdleConns   = 10
	defaultMaxOpenConns   = 100
)

var ctx context.Context

func Config(c context.Context, conf embed.FS) {
	ctx = c
	box := models.ConfBox{
		Ctx: c,
		Fs:  conf,
		Dir: configDir,
	}

	v := viper.New()
	readConfig(box, v, configFile)
	settings := v.AllSettings()
	for i, s := range settings {
		v.SetDefault(i, s)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		log.Panic(fmt.Sprintf("Config 初始化失败：%s", err.Error()))
	}

	if global.Config.Server.ConnectTimeout < 1 {
		global.Config.Server.ConnectTimeout = defaultConnectTimeout
	}

	if strings.TrimSpace(global.Config.Server.ApiPrefix) == "" {
		global.Config.Server.ApiPrefix = "api"
	}

	if strings.TrimSpace(global.Config.Server.ApiVersion) == "" {
		global.Config.Server.ApiVersion = "v1"
	}

	if global.Config.Mysql.MaxIdleConns < 1 {
		global.Config.Mysql.MaxIdleConns = defaultMaxIdleConns
	}

	if global.Config.Mysql.MaxOpenConns < 1 {
		global.Config.Mysql.MaxOpenConns = defaultMaxOpenConns
	}

	log.Info("Config 初始化成功...")
}

func readConfig(box models.ConfBox, v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := box.Get(configFile)
	if len(config) == 0 {
		log.Panic(fmt.Sprintf("Config 初始化失败, 配置文件未找到，路径：%s", box.Dir))
	}
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		log.Panic(fmt.Sprintf("Config 初始化失败, 配置文件未找到，路径：%s", box.Dir))
	}
}
