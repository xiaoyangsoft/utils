//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 8:03 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package sms

// 发送短信
func Send(mobile, template, payload string) (result *Result, err error) {
	return client().Send(mobile, template, payload)
}
