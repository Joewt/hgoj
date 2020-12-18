package redis

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/garyburd/redigo/redis"
	"time"
)


var (
	pool *redis.Pool
	redisHost = beego.AppConfig.String("redishost")+":"+beego.AppConfig.String("redisport")
	redisPass = beego.AppConfig.String("redispwd")
)

func init() {
	pool = newRedisPool()
}

func RedisPool() *redis.Pool {
	return pool
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				logs.Error("redis dial err:"+err.Error())
				return nil, err
			}

			if _, err := c.Do("AUTH", redisPass); err != nil {
				logs.Error("redis密码验证失败")
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:      2000,
		MaxActive:    1000,
		IdleTimeout:  time.Second * 300,
		Wait:         false,
	}
}



