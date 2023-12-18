package gredis

import (
	"errors"
	"github.com/go-redis/redis/v8"
)

// InitRedis
/**
 * @Description: 初始化Redis连接池
 * @param ip IP地址
 * @param port 端口号
 * @param password 密码
 * @param db 数据库序号
 * @return rdb 数据库连接池
 * @return err 错误
 */
func InitRedis(ip, port, password string, db int) (rdb *redis.Client, err error) {
	// 初始化连接池
	rdb = redis.NewClient(&redis.Options{
		Addr:     ip + ":" + port,
		Password: password,
		DB:       db,
	})
	// 使用Ping检测检测状态
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		err = errors.New("ping err :" + err.Error())
		return
	}
	RDB = rdb
	return
}
