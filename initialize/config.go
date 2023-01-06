package initialize

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/spf13/viper"
	"strings"
)

const (
	configType            = "yaml"
	debugConfig           = "conf/config.debug.yml"
	testConfig            = "conf/config.test.yml"
	releaseConfig         = "conf/config.release.yml"
	defaultConnectTimeout = 5
)

var ctx context.Context

func Config(c context.Context, conf embed.FS) {
	var config string

	ctx = c
	switch global.ProjectEnv {
	case "debug":
		config = debugConfig
	case "release":
		config = releaseConfig
	default:
		config = debugConfig
	}

	fmt.Printf("项目环境为：%s，配置文件为：%s \n", global.ProjectEnv, config)

	v := viper.New()
	v.SetConfigType(configType)
	v.SetConfigFile(config)
	if err := v.ReadConfig(bytes.NewReader(readConfig(conf, config))); err != nil {
		panic(fmt.Errorf("读取配置文件失败：%s", err.Error()))
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		panic(errors.Wrapf(err, "初始化配置文件失败"))
	}

	if global.Config.Server.ConnectTimeout < 1 {
		global.Config.Server.ConnectTimeout = defaultConnectTimeout
	}

	if strings.TrimSpace(global.Config.Server.ApiVersion) == "" {
		global.Config.Server.ApiVersion = "v1"
	}
}

func readConfig(fs embed.FS, config string) (bs []byte) {
	var err error
	bs, err = fs.ReadFile(config)
	if err != nil {
		fmt.Printf("读取文件错误,err:%s \n", err.Error())
	}
	return
}
