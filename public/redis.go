package public

import (
	"github.com/go-redis/redis"
	"time"
)

var RedisDb *redis.Client

func InitClirnt() (err error) {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "101.43.6.142:6379",
		Password: "123456",
		DB:       0,
	})
	_, err = RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func RedisSet(key string, value string, time time.Duration) (err error) {
	err = RedisDb.Set(key, value, time).Err()

	if err != nil {
		return err
	}
	return nil
}
func RedisGet(key string) (val string, err error) {
	val, err = RedisDb.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
