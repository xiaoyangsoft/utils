//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:41 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package string

import "unsafe"

// StrToBytes 字符串转字节数组
func StrToBytes(v string) []byte {
	return *(*[]byte)(unsafe.Pointer(&v))
}
