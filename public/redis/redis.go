package redis

import (
	"gindemo01/common"
	"gindemo01/public"
	"github.com/go-redis/redis"
	"time"
)

var RedisDb *redis.Client

func InitClirnt() (err error) {
	addr := common.Redisinfo.Host + ":" + common.Redisinfo.Port
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: common.Redisinfo.Passwd,
		DB:       0,
	})
	_, err = RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Set(key string, value string, time time.Duration) (err error) {
	err = RedisDb.Set(key, value, time).Err()

	if err != nil {
		return err
	}
	return nil
}
func Get(key string) (val string, err error) {
	val, err = RedisDb.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
func SetCaptcha(e string) string {
	caps := public.Captcha()
	err := RedisDb.Set(caps, e, time.Minute).Err()
	if err != nil {
		return "err"
	}
	return caps
}
func GetCaptcha(c string) string {
	val, err := RedisDb.Get(c).Result()
	if err != nil {
		return "err"
	}
	return val
}
func Del(key string) {
	RedisDb.Del(key)
}
