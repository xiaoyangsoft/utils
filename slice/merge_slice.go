//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:41 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

// MergeSlice 将多个slice合并成一个slice
func MergeSlice[T any](slices ...[]T) []T {
	var newSlice []T
	for _, slice := range slices {
		for _, v := range slice {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
