//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:43 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

import "math"

// Ceil 对float数据向上取整
func Ceil[T float32 | float64](num T) int {
	return int(math.Ceil(float64(num)))
}