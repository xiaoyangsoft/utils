//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:45 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

// InSlice 判断value是否在slice中
func InSlice[T comparable](value T, slices []T) bool {
	for _, v := range slices {
		if value == v {
			return true
		}
	}
	return false
}
