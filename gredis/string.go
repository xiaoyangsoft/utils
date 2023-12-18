package gredis

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

// Set
/**
 * @Description: 设置指定 key 的值
 * @param key Key名
 * @param value 值
 * @param expiration 有效时间（不过期写0）
 * @return err 错误
 */
func Set(key string, value interface{}, expiration time.Duration) (err error) {
	err = RDB.Set(ctx, key, value, expiration).Err()
	return
}

// MSet
/**
 * @Description: 批量设置指定 key 的值
 * @param data
 * @return err
 */
func MSet(data map[string]interface{}) (err error) {
	if len(data) == 0 {
		return
	}
	values := make([]interface{}, len(data)*2)
	i := 0
	for k, v := range data {
		values[i] = k
		values[i+1] = v
		i += 2
	}
	err = RDB.MSet(ctx, values).Err()
	return
}

// Get
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func Get(key string) (value string, err error) {
	value, err = RDB.Get(ctx, key).Result()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// MGet
/**
 * @Description: 批量获取 key 对应的值
 * @param keys
 * @return values
 * @return err
 */
func MGet(keys ...string) (values []interface{}, err error) {
	values, err = RDB.MGet(ctx, keys...).Result()
	return
}

// GetInt64
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func GetInt64(key string) (value int64, err error) {
	value, err = RDB.Get(ctx, key).Int64()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// GetInt
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func GetInt(key string) (value int, err error) {
	value, err = RDB.Get(ctx, key).Int()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// GetUint64
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func GetUint64(key string) (value uint64, err error) {
	value, err = RDB.Get(ctx, key).Uint64()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// GetBool
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func GetBool(key string) (value bool, err error) {
	value, err = RDB.Get(ctx, key).Bool()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// GetFloat32
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func GetFloat32(key string) (value float32, err error) {
	value, err = RDB.Get(ctx, key).Float32()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// GetFloat
/**
 * @Description: 获取指定 key 的值
 * @param key 键名
 * @return value
 * @return err
 */
func GetFloat(key string) (value float64, err error) {
	value, err = RDB.Get(ctx, key).Float64()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return
	} else if err != nil {
		return
	} else {
		return
	}
}

// Incr
/**
 * @Description: 自增1
 * @param key
 * @return err
 */
func Incr(key string) (err error) {
	err = RDB.Incr(ctx, key).Err()
	return
}

// IncrBy
/**
 * @Description: 自增指定步长（整型）
 * @param key 键名
 * @param value 步长
 * @return err
 */
func IncrBy(key string, value int64) (err error) {
	err = RDB.IncrBy(ctx, key, value).Err()
	return
}

// IncrByFloat
/**
 * @Description: 自增指定步长(浮点型)
 * @param key 键名
 * @param value 步长
 * @return err
 */
func IncrByFloat(key string, value float64) (err error) {
	err = RDB.IncrByFloat(ctx, key, value).Err()
	return
}
