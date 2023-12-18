//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:39 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package valid

import (
	"regexp"
	"unicode"
)

// IsAllChinese 验证给定的字符串全部为中文
func IsAllChinese(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Scripts["Han"], r) {
			return false
		}
	}
	return true
}

// IsContainChinese 验证给定的字符串包含中文
func IsContainChinese(input string) bool {
	for _, r := range input {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

// IsChineseName 验证是否为中文名
func IsChineseName(name string) bool {
	pattern := "^[\u4E00-\u9FA5]{2,6}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(name)
}

// IsEnglishName 验证是否为英文名
func IsEnglishName(name string) bool {
	match, _ := regexp.MatchString(`^([a-zA-Z]+\s)*[a-zA-Z]+$`, name)
	return match
}
