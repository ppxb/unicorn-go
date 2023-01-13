package initialize

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ppxb/unicorn/pkg/global"
	"github.com/ppxb/unicorn/pkg/query"
	"strings"
	"time"
)

func Redis() {
	if strings.TrimSpace(global.Config.Redis.Uri) == "" {
		return
	}

	init := false
	ctx, cancel := context.WithTimeout(ctx, time.Duration(global.Config.Server.ConnectTimeout)*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				if !init {
					panic(fmt.Sprintf("[初始化] Redis失败: 连接超时(%ds)", global.Config.Server.ConnectTimeout))
				}
				return
			}
		}
	}()

	client, err := query.ParseRedisUri(global.Config.Redis.Uri)
	if err != nil {
		panic(errors.Wrap(err, "[初始化] Redis失败"))
	}

	err = client.Ping(ctx).Err()
	if err != nil {
		panic(errors.Wrap(err, "[初始化] Redis失败"))
	}
	global.Redis = client

	init = true
	fmt.Println("[初始化] Redis成功")
}
