//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:40 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package valid

import "regexp"

// IsMobile 验证是否为手机号码
func IsMobile(phone string) bool {
	match, _ := regexp.MatchString(`^1[3456789]\d{9}$`, phone)
	return match
}

// IsTelephone 验证是否为座机号码
func IsTelephone(telephone string) bool {
	match, _ := regexp.MatchString(`^0\d{2,3}-?\d{7,8}$`, telephone)
	return match
}
