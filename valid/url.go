//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:41 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package valid

import "regexp"

// IsURL 是否为URL地址
func IsURL(url string) bool {
	match, _ := regexp.MatchString(`^(http|https)://[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?$`, url)
	return match
}
