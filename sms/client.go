//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 8:00 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package sms

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/tidwall/gjson"
	"sync"
)

var (
	initClient *Client
	onceClient sync.Once
)

type Client struct {
	opts *options
}

func NewClient(opts ...Option) *Client {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	onceClient.Do(func() {
		initClient = &Client{
			opts: &o,
		}
	})
	return initClient
}

func client() *Client {
	if initClient == nil {
		panic("sms client is not initialized")
	}
	return initClient
}

// 发送短信
func (c *Client) Send(mobile, template, payload string) (result *Result, err error) {
	smsRequest := dysmsapi.CreateSendSmsRequest()
	smsRequest.SignName = c.opts.signName
	smsRequest.PhoneNumbers = mobile   //手机号
	smsRequest.TemplateCode = template //模版编码
	smsRequest.TemplateParam = payload //附加参数
	client, err := dysmsapi.NewClientWithAccessKey(c.opts.regionId, c.opts.accessKeyId, c.opts.accessKeySecret)
	if err != nil {
		return nil, err
	}
	response, err := client.SendSms(smsRequest)
	if err != nil {
		return nil, err
	}
	result = &Result{
		RequestId: response.RequestId,
		BizId:     response.BizId,
		Code:      response.Code,
		Message:   response.Message,
	}
	bytes := response.GetHttpContentBytes()
	responseCode := gjson.GetBytes(bytes, "Code").String()
	if response.IsSuccess() && responseCode == "OK" {
		return result, nil
	}
	return result, errors.New(gjson.GetBytes(bytes, "Message").String())
}

// result 响应结果
type Result struct {
	RequestId string `json:"request_id"`
	BizId     string `json:"biz_id"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

func (r *Result) String() string {
	buf, _ := json.Marshal(r)
	return string(buf)
}

// 获取RequestId
func (r *Result) GetRequestId() string {
	return r.RequestId
}

// 获取BizId
func (r *Result) GetBizId() string {
	return r.BizId
}

// 获取Code
func (r *Result) GetCode() string {
	return r.Code
}

// 获取Message
func (r *Result) GetMessage() string {
	return r.Message
}
