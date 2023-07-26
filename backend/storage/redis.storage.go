package storage

import (
	"context"
	"fmt"
	"sync"
	"time"

	"git.eon-cds.de/repos/dlab/wad-fido2/backend/utils"
	"github.com/redis/go-redis/v9"
)

type RedisStorageInterface interface {
	SetWithTTL(string, string, time.Duration) error
	Get(string) (string, error)
}

type RedisStorage struct {
	rc *redis.Client
}

var (
	redisSyncOnce sync.Once
	redisIntance  RedisStorageInterface
)

func GetRedisClientInstance() RedisStorageInterface {

	envConfig := utils.ParseEnv()

	redisSyncOnce.Do(func() {
		rc := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", envConfig.Redis_Host, envConfig.Redis_Port),
			Password: envConfig.Redis_Password,
			DB:       0,
		})

		redisIntance = &RedisStorage{
			rc: rc,
		}
	})

	return redisIntance
}

func (rs *RedisStorage) SetWithTTL(key string, data string, ttl time.Duration) error {

	_, err := rs.rc.SetNX(context.Background(), key, data, ttl).Result()
	return err

}

func (rs *RedisStorage) Get(key string) (string, error) {
	return rs.rc.Get(context.Background(), key).Result()

}
