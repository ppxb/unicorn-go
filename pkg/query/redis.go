package query

import (
	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

func ParseRedisUri(uri string) (client redis.UniversalClient, err error) {
	var opt asynq.RedisConnOpt
	if uri != "" {
		opt, err = asynq.ParseRedisURI(uri)
		if err != nil {
			return
		}
		client = opt.MakeRedisClient().(redis.UniversalClient)
		return
	}
	err = errors.Errorf("invalid redis config")
	return
}
