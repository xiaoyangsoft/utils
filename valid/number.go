//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:40 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package valid

import "regexp"

// IsNumber 验证是否全部为数字
func IsNumber(input string) bool {
	reg := regexp.MustCompile("^[0-9]+$")
	return reg.MatchString(input)
}
