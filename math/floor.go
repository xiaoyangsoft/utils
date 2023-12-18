//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:44 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

import "math"

// Floor 对float数据向下取整
func Floor[T float32 | float64](num T) int {
	return int(math.Floor(float64(num)))
}
