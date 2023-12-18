//============================================================
// 作者: 杨大雷
// 日期: 2023/12/02 4:45 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

import "math"

// Round 对float数据四舍五入
func Round[T float32 | float64](num T) int {
	return int(math.Round(float64(num)))
}
