//============================================================
// 描述: todo
// 作者: 杨大雷
// 日期: 2023/12/02 8:30 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package oss

import (
	"context"
	"errors"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
)

type Tencent struct {
	client *cos.Client
	config *Config
}

func NewTencent(config Config) (*Tencent, error) {
	tencent := Tencent{}
	err := tencent.init(config)
	if err != nil {
		return nil, err
	}
	return &tencent, err
}

func (t *Tencent) init(config Config) error {
	if t.config == nil {
		t.config = &config
	}

	if config.AppId == "" || config.Region == "" {
		return errors.New("configuration not correct")
	}

	if t.client == nil {
		c := config
		buckerURL, _ := url.Parse(fmt.Sprintf("http://%s-%s.cos.%s.myqcloud.com", c.Bucket, c.AppId, c.Region))

		client := cos.NewClient(&cos.BaseURL{BucketURL: buckerURL}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  c.KeyId,
				SecretKey: c.KeySecret,
			},
		})
		t.client = client
	}

	return nil
}

// 获取文件
func (t *Tencent) GetObjectToFile(objectKey, filePath string) error {
	response, err := t.client.Object.Download(context.Background(), objectKey, filePath, nil)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return httpError(response)
	}

	return nil
}

// 删除文件
func (t *Tencent) DeleteObject(objectKey string) error {
	response, err := t.client.Object.Delete(context.Background(), objectKey)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return httpError(response)
	}

	return nil
}

func (t *Tencent) PutObject(objectKey string, reader io.Reader) error {
	response, err := t.client.Object.Put(context.Background(), objectKey, reader, nil)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return httpError(response)
	}

	return nil
}

// PutObjectFromFile 上传文件
func (t *Tencent) PutObjectFromFile(objectKey, filePath string) error {
	response, err := t.client.Object.PutFromFile(context.Background(), objectKey, filePath, nil)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return httpError(response)
	}

	return nil
}

func (t *Tencent) IsExists(objectKey string) (bool, error) {
	return t.client.Object.IsExist(context.Background(), objectKey)
}

// cos 请求错误
func httpError(response *cos.Response) error {
	bytes, err := io.ReadAll(response.Body)
	defer func() {
		err = response.Body.Close()
	}()
	if err != nil {
		return err
	}

	return errors.New(string(bytes))
}

// GetObjectUrl 获取访问URL
func (t *Tencent) GetObjectUrl(objectKey string) string {
	if t.config.Domain == "" {
		return fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com/%s",
			t.config.Bucket,
			t.config.AppId,
			t.config.Region,
			objectKey,
		)
	}
	return fmt.Sprintf("https://%s/%s", t.config.Domain, objectKey)
}
