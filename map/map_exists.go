//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:47 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package _map

// MapKeyExists 判断map中的key是否存在
func MapKeyExists[T comparable, T2 any](m map[T]T2, key T) bool {
	_, exists := m[key]
	return exists
}

// MapValueExists 判断map中的value是否存在
func MapValueExists[T comparable, T2 comparable](m map[T2]T, value T) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}
