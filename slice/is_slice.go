//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:46 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

import "reflect"

// IsSlice 判断指定值i是否是slice类型
func IsSlice(slice any) bool {
	return reflect.ValueOf(slice).Kind() == reflect.Slice
}
