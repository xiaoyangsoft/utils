//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:44 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

// Min 返回slice中最小值
func Min[T Numeric](slice []T) T {
	if len(slice) == 0 {
		return 0
	}
	min := slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
	}
	return min
}
