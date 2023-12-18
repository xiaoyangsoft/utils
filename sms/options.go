//============================================================
// 作者: 杨大雷
// 日期: 2023/12/02 7:59 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package sms

var defaultOptions = options{
	signName: "晓杨信息科技",
}

type options struct {
	regionId        string // 区域
	accessKeyId     string // key
	accessKeySecret string // 秘钥
	signName        string // 签名
}

// Option 配置项
type Option func(o *options)

// 设置区域
func SetRegionId(regionId string) Option {
	return func(o *options) {
		o.regionId = regionId
	}
}

// 设置 KEY
func SetAccessKeyId(accessKeyId string) Option {
	return func(o *options) {
		o.accessKeyId = accessKeyId
	}
}

// 设置秘钥
func SetAccessKeySecret(accessKeySecret string) Option {
	return func(o *options) {
		o.accessKeySecret = accessKeySecret
	}
}

// 设置签名
func SetSignName(signName string) Option {
	return func(o *options) {
		o.signName = signName
	}
}
