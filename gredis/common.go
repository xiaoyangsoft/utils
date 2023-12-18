package gredis

import (
	"errors"
	"time"
)

// Keys
/**
 * @Description: 获取所有Keys
 * @param pattern
 * @return keys
 * @return err
 */
func Keys(pattern string) (keys []string, err error) {
	keys, err = RDB.Keys(ctx, pattern).Result()
	return
}

// KeysByScan
/**
 * @Description: 利用Scan获取Keys（生产环境用于替换Keys）
 * @param cursor 游标
 * @param match 查询条件
 * @param count 每次取出的条数
 * @return keys Keys
 * @return err 错误
 */
func KeysByScan(cursor uint64, match string, count int64) (keys []string, err error) {
	keys, cursor, err = RDB.Scan(ctx, cursor, match, count).Result()
	if err != nil {
		err = errors.New("scan keys failed err:" + err.Error())
		return
	}
	return
}

// Size
/**
 * @Description: 获取Keys的数量
 * @return size
 * @return err
 */
func Size() (size int64, err error) {
	size, err = RDB.DBSize(ctx).Result()
	return
}

// Delete
/**
 * @Description: 删除键
 * @param keys
 * @return err
 */
func Delete(keys ...string) (err error) {
	err = RDB.Del(ctx, keys...).Err()
	return
}

// Exists
/**
 * @Description: 判断键是否存在
 * @param key
 * @return exists
 * @return err
 */
func Exists(key string) (exists bool, err error) {
	var result int64
	result, err = RDB.Exists(ctx, key).Result()
	if err != nil {
		return
	}
	if result > 0 {
		exists = true
	} else {
		exists = false
	}
	return
}

// Expire
/**
 * @Description: 设置Key的有效期
 * @param key 键名
 * @param expiration 有效期
 * @return err 错误
 */
func Expire(key string, expiration time.Duration) (err error) {
	err = RDB.Expire(ctx, key, expiration).Err()
	return
}

// TTL
/**
 * @Description: 查看Key剩余的有效期
 * @param key 键名
 * @return expiration 有效期
 * @return err 错误
 */
func TTL(key string) (expiration time.Duration, err error) {
	expiration, err = RDB.TTL(ctx, key).Result()
	return
}
