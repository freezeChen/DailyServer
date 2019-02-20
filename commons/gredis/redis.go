package gredis

import (
	"DailyServer/commons/config"
	"DailyServer/commons/glog"
	"github.com/garyburd/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

func InitRedis() error {
	RedisConn = &redis.Pool{
		MaxIdle:     1024,
		MaxActive:   60000,
		IdleTimeout: 200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.DefaultConfig.RedisConn,
				redis.DialPassword(config.DefaultConfig.Redispwd),
				redis.DialConnectTimeout(200*time.Millisecond),
				redis.DialReadTimeout(500*time.Millisecond),
				redis.DialWriteTimeout(500*time.Millisecond),
			)

			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()



	_, err := conn.Do("SET", key, data)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func HSet(key string, item string, value string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	if err := conn.Send("HSET", key, item, value); err != nil {
		glog.Errorf("conn HSET (%s %s %s)", key, item, value)
		return err
	}

	return nil
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
