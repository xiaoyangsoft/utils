//============================================================
// 描述: todo
// 作者: 杨大雷
// 日期: 2023/12/02 8:21 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

type AliYun struct {
	config *Config
	bucket *oss.Bucket
}

func NewAliYun(config Config) (*AliYun, error) {
	aliYun := AliYun{}
	err := aliYun.init(config)
	if err != nil {
		return nil, err
	}
	return &aliYun, nil
}

func (aliYun *AliYun) init(config Config) error {
	if aliYun.config == nil {
		aliYun.config = &config
	}
	if config.Endpoint == "" {
		return errors.New("AliYun configuration Endpoint is not correct")
	}
	if aliYun.bucket == nil {
		client, err := oss.New(config.Endpoint, config.KeyId, config.KeySecret)
		if err != nil {
			return err
		}
		bucket, err := client.Bucket(config.Bucket)
		if err != nil {
			return err
		}
		aliYun.bucket = bucket
	}
	return nil
}

// GetObjectToFile 获取文件
func (aliYun *AliYun) GetObjectToFile(objectKey, filePath string) error {
	return aliYun.bucket.GetObjectToFile(objectKey, filePath)
}

// DeleteObject 删除文件
func (aliYun *AliYun) DeleteObject(objectKey string) error {
	return aliYun.bucket.DeleteObject(objectKey)
}

func (aliYun *AliYun) PutObject(objectKey string, reader io.Reader) error {
	return aliYun.bucket.PutObject(objectKey, reader)
}

func (aliYun *AliYun) PutObjectFromFile(objectKey, filePath string) error {
	return aliYun.bucket.PutObjectFromFile(objectKey, filePath)
}

func (aliYun *AliYun) IsExists(objectKey string) (bool, error) {
	return aliYun.bucket.IsObjectExist(objectKey)
}

// GetObjectUrl 获取访问URL
func (aliYun *AliYun) GetObjectUrl(objectKey string) string {
	if aliYun.config.Domain == "" {
		return fmt.Sprintf("https://%s.%s/%s",
			aliYun.config.Bucket,
			aliYun.config.Endpoint,
			objectKey,
		)
	}
	return fmt.Sprintf("https://%s/%s", aliYun.config.Domain, objectKey)
}
