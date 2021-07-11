package redisclient

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

//redis连接池

var pool *redis.Pool

func init() {
	// 从配置文件获取redis的ip以及db
	REDIS_HOST := "localhost"
	REDIS_PORT := 6379
	REDIS_DB := 0
	// 建立连接池
	pool = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,                 //最大空闲连接数量，即使没有redis连接时依然可以保持N个空闲的连接
		MaxActive:   10,                //最多同时激活的连接数
		IdleTimeout: 180 * time.Second, //空闲连接等待时间
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", REDIS_HOST, REDIS_PORT))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
}
func Get() redis.Conn {
	return pool.Get()
}
