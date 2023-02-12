package redis

import (
	"gindemo01/common"
	"gindemo01/public"
	"github.com/go-redis/redis"
	"time"
)

var redisDb0 *redis.Client
var redisDb1 *redis.Client
var redisDb2 *redis.Client
var redisDb3 *redis.Client
var redisDb4 *redis.Client

func InitClirnt() (err error) {
	addr := common.Redisinfo.Host + ":" + common.Redisinfo.Port
	redisDb0 = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: common.Redisinfo.Passwd,
		DB:       0,
	})
	_, err = redisDb0.Ping().Result()
	redisDb1 = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: common.Redisinfo.Passwd,
		DB:       1,
	})
	_, err = redisDb1.Ping().Result()
	redisDb2 = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: common.Redisinfo.Passwd,
		DB:       2,
	})
	_, err = redisDb2.Ping().Result()
	redisDb3 = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: common.Redisinfo.Passwd,
		DB:       3,
	})
	_, err = redisDb3.Ping().Result()
	redisDb4 = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: common.Redisinfo.Passwd,
		DB:       4,
	})
	_, err = redisDb4.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Set(key string, value string, time time.Duration, d int) (err error) {
	var redisDb *redis.Client
	switch d {
	case 0:
		redisDb = redisDb0
	case 1:
		redisDb = redisDb1
	case 2:
		redisDb = redisDb2
	case 3:
		redisDb = redisDb3
	case 4:
		redisDb = redisDb4
	default:
		redisDb = redisDb0
	}
	err = redisDb.Set(key, value, time).Err()
	if err != nil {
		return err
	}
	return nil
}
func Get(key string, d int) (val string, err error) {
	var redisDb *redis.Client
	switch d {
	case 0:
		redisDb = redisDb0
	case 1:
		redisDb = redisDb1
	case 2:
		redisDb = redisDb2
	case 3:
		redisDb = redisDb3
	case 4:
		redisDb = redisDb4
	default:
		redisDb = redisDb0
	}
	val, err = redisDb.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
func SetCaptcha(e string, time time.Duration, d int) string {
	var redisDb *redis.Client
	switch d {
	case 0:
		redisDb = redisDb0
	case 1:
		redisDb = redisDb1
	case 2:
		redisDb = redisDb2
	case 3:
		redisDb = redisDb3
	case 4:
		redisDb = redisDb4
	default:
		redisDb = redisDb0
	}
	caps := public.Captcha()
	err := redisDb.Set(caps, e, time).Err()
	if err != nil {
		return "err"
	}
	return caps
}
func GetCaptcha(c string, d int) (string, error) {
	var redisDb *redis.Client
	switch d {
	case 0:
		redisDb = redisDb0
	case 1:
		redisDb = redisDb1
	case 2:
		redisDb = redisDb2
	case 3:
		redisDb = redisDb3
	case 4:
		redisDb = redisDb4
	default:
		redisDb = redisDb0
	}
	val, err := redisDb.Get(c).Result()
	return val, err
}
func Del(key string, d int) {
	var redisDb *redis.Client
	switch d {
	case 0:
		redisDb = redisDb0
	case 1:
		redisDb = redisDb1
	case 2:
		redisDb = redisDb2
	case 3:
		redisDb = redisDb3
	case 4:
		redisDb = redisDb4
	default:
		redisDb = redisDb0
	}
	redisDb.Del(key)
}
