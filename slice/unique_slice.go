//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:42 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

// UniqueSlice 移除slice中重复的值
func UniqueSlice[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	m := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}
