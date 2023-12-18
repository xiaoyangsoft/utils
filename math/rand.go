//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:44 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

import (
	"math/rand"
	"time"
)

// Rand 生成随机整数
func Rand[T Numeric](min, max T) T {
	rand.Seed(time.Now().UnixNano())
	return T(rand.Intn(int(max)-int(min)+1) + int(min))
}
