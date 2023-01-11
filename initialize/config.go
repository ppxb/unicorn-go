package initialize

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/model"
	"github.com/spf13/viper"
	"strings"
)

const (
	configType            = "yaml"
	configDir             = "conf"
	debugConfig           = "config.debug.yml"
	testConfig            = "config.test.yml"
	releaseConfig         = "config.release.yml"
	defaultConnectTimeout = 5
)

var ctx context.Context

func Config(c context.Context, conf embed.FS) {
	var configFile string
	var box model.ConfBox

	ctx = c
	box.Ctx = ctx
	box.Fs = conf
	box.Dir = configDir
	global.ConfBox = box

	switch gin.Mode() {
	case gin.TestMode:
		configFile = testConfig
	case gin.ReleaseMode:
		configFile = releaseConfig
	default:
		configFile = debugConfig
	}

	v := viper.New()
	readConfig(box, v, configFile)
	settings := v.AllSettings()
	for i, s := range settings {
		v.SetDefault(i, s)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(errors.Wrapf(err, "初始化Config失败"))
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

	fmt.Println("初始化Config成功")
}

func readConfig(box model.ConfBox, v *viper.Viper, configFile string) {
	v.SetConfigType(configType)
	config := box.Get(configFile)
	if len(config) == 0 {
		panic(fmt.Sprintf("初始化Config失败, 配置文件路径：%s", box.Dir))
	}
	if err := v.ReadConfig(bytes.NewReader(config)); err != nil {
		panic(errors.Wrapf(err, "初始化Config失败, 配置文件路径：%s`", box.Dir))
	}
}
