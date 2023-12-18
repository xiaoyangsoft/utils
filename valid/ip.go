//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:39 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package valid

import "net"

// IsIPv4 是否为ipv4地址
func IsIPv4(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() != nil
}

// IsIPv6 是否为ipv6地址
func IsIPv6(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() == nil
}
