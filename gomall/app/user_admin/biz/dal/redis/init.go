package redis

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/redis/go-redis/v9"
	"github.com/yqihe/91-mall/gomall/app/user_admin/conf"
)

var (
	RedisClient *redis.Client
)

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func GetDataFromCache(ctx context.Context, key string) ([]byte, error) {
	klog.Debugf("getDataFromCache, key: %s", key)
	return RedisClient.Get(ctx, key).Bytes()
}

func SetCache(ctx context.Context, key string, data []byte) error {
	klog.Debugf("setCache, key: %s", key)
	return RedisClient.Set(ctx, key, data, 0).Err()
}
