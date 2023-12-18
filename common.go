//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:28 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package utils

// 三目运算
func Ternary[T any](condition bool, one, two T) T {
	if condition {
		return one
	}
	return two
}
