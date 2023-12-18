//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:41 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// SumSlice 对slice中的元素求和
func SumSlice[T Numeric](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
