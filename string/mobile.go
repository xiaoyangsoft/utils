//============================================================
// 描述: todo
// 作者: 杨前磊
// 日期: 2023/12/02 8:38 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package string

// 手机号中间4位替换为*号
func FormatMobileStar(mobile string) string {
	if len(mobile) <= 10 {
		return mobile
	}
	return mobile[:3] + "****" + mobile[7:]
}
