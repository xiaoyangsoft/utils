//============================================================
// 描述: todo
// 作者: 杨大雷
// 日期: 2023/12/02 8:28 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package oss

import (
	"errors"
	"fmt"
	"github.com/minio/minio-go"
	"io"
)

type Minio struct {
	config *Config
	client *minio.Client
	bucket string
}

func NewMinio(config Config) (*Minio, error) {
	mini := Minio{}
	err := mini.init(config)
	if err != nil {
		return nil, err
	}
	return &mini, nil
}

func (mini *Minio) init(config Config) error {
	if mini.config == nil {
		mini.config = &config
	}
	if config.Endpoint == "" {
		return errors.New("AliYun configuration Endpoint is not correct")
	}
	if mini.client == nil {
		client, err := minio.New(config.Endpoint, config.KeyId, config.KeySecret, config.UseSSL)
		if err != nil {
			return errors.New("minio Collection Failed")
		}
		mini.client = client
	}
	if mini.bucket == "" {
		if config.Bucket == "" {
			return errors.New("bucket Not specified")
		}
		mini.bucket = config.Bucket
	}
	return nil
}

// GetObjectToFile 获取文件
func (mini *Minio) GetObjectToFile(objectKey, filePath string) error {
	return mini.client.FGetObject(mini.bucket, objectKey, filePath, minio.GetObjectOptions{})
}

// DeleteObject 删除文件
func (mini *Minio) DeleteObject(objectKey string) error {
	return mini.client.RemoveObject(mini.bucket, objectKey)
}

func (mini *Minio) PutObject(objectKey string, reader io.Reader) error {
	return errors.New("method is not support")
}

func (mini *Minio) PutObjectFromFile(objectKey, filePath string) error {
	_, err := mini.client.FPutObject(mini.bucket, objectKey, filePath, minio.PutObjectOptions{})
	return err
}

func (mini *Minio) IsExists(objectKey string) (bool, error) {
	return false, errors.New("method is not support")
}

// GetObjectUrl 获取访问URL
func (mini *Minio) GetObjectUrl(objectKey string) string {
	if mini.config.Domain == "" {
		return fmt.Sprintf("http://%s/%s/%s",
			mini.config.Endpoint,
			mini.config.Bucket,
			objectKey,
		)
	}
	return fmt.Sprintf("https://%s/%s", mini.config.Domain, objectKey)
}
