package model

//定义了操作redis的基本函数
import (
	redisclient "github.com/Jundong-chan/admin/authorize/pkg/redisconfig"
	"errors"

	Redis "github.com/gomodule/redigo/redis"
)

type RedisService struct {
}

//存放数据
func (redis RedisService) Set(key string, value interface{}) error {
	c := redisclient.Get()
	defer c.Close()
	_, err := c.Do("SET", key, value)
	if err != nil {
		return errors.New("setting in redis error:" + err.Error())
	}
	return nil
}

//查询key是否存在
func (redis RedisService) IsExist(key string) (bool, error) {
	c := redisclient.Get()
	defer c.Close()
	exists, err := Redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return exists, nil
}

//根据key返回内容
func (redis RedisService) Get(key string) (value interface{}, err error) {
	c := redisclient.Get()
	defer c.Close()
	value, err = Redis.String(c.Do("GET", key))
	if err != nil {
		return nil, errors.New("getting in redis error:" + err.Error())
	}
	return
}

func (redis RedisService) Delete(key string) error {
	c := redisclient.Get()
	defer c.Close()
	_, err := c.Do("DEL", key)
	if err != nil {
		return errors.New("delete in redis error:" + err.Error())
	}
	return nil
}
