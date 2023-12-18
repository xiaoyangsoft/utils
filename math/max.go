//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:44 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

// Max 返回slice中最大值
func Max[T Numeric](slice []T) T {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
