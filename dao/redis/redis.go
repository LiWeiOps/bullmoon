package redis

import (
	"bullmoon/settings"
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init(conf *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,  // 密码
		DB:       conf.Db,        // 数据库
		PoolSize: conf.Pool_size, // 连接池大小
	})
	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis conn failed err:", zap.Error(err))
		return err
	}
	return
}

func Close() {
	rdb.Close()
}
