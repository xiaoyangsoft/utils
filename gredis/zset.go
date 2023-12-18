package gredis

import "github.com/go-redis/redis/v8"

// ZAdd
/**
 *  @Description: 向有序集合添加一个元素
 *  @param key 集合名称
 *  @param score
 *  @param member
 *  @return err
 */
func ZAdd(key string, score float64, member string) (err error) {
	err = RDB.ZAdd(ctx, key, &redis.Z{
		Score:  score,
		Member: member,
	}).Err()
	return
}

// ZCard
/**
 *  @Description: 返回集合元素个数
 *  @param key 集合名称
 *  @return result 元素个数
 *  @return err
 */
func ZCard(key string) (result int64, err error) {
	result, err = RDB.ZCard(ctx, key).Result()
	return
}

// ZRem
/**
 *  @Description: 删除集合元素
 *  @param key 集合名称
 *  @param members 元素列表
 *  @return err
 */
func ZRem(key string, members []string) (err error) {
	err = RDB.ZRem(ctx, key, members).Err()
	return
}

// ZRange
/**
 *  @Description: 返回指的范围的元素
 *  @param key 集合名称
 *  @param start 起始位置
 *  @param end 结束位置
 *  @return result
 */
func ZRange(key string, start, end int64) (result []string, err error) {
	result, err = RDB.ZRange(ctx, key, start, end).Result()
	return
}

// ZLexCount
/**
 *  @Description: 计算指定字典区间内成员数量
 *  @param key 集合名称
 *  @param min
 *  @param max
 *  @return result
 *  @return err
 */
func ZLexCount(key string, min, max string) (result int64, err error) {
	result, err = RDB.ZLexCount(ctx, key, min, max).Result()
	return
}

// ZCount
/**
 *  @Description: 指定区间分数的成员数
 *  @param key 集合名称
 *  @param min
 *  @param max
 *  @return result
 *  @return err
 */
func ZCount(key string, min, max string) (result int64, err error) {
	result, err = RDB.ZCount(ctx, key, min, max).Result()
	return
}

// ZRangeByScore
/**
 *  @Description: 查询范围内容的设备
 *  @param key
 *  @param min
 *  @param max
 *  @return result
 *  @return err
 */
func ZRangeByScore(key string, min, max string) (result []string, err error) {
	result, err = RDB.ZRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: min,
		Max: max,
	}).Result()
	return
}

// ZRemRangeByScore
/**
 *  @Description: 删除指定分数之间的元素
 *  @param key
 *  @param min
 *  @param max
 *  @return err
 */
func ZRemRangeByScore(key string, min, max string) (err error) {
	err = RDB.ZRemRangeByScore(ctx, key, min, max).Err()
	return
}
