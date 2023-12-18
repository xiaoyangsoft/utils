//============================================================
// 描述: todo
// 作者: 杨大雷
// 日期: 2023/12/02 8:30 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package oss

import (
	"errors"
	"io"
)

const (
	AliYunConst  = "aliYun"
	TencentConst = "tencent"
	QinNiuConst  = "qiNiu"
	MinioConst   = "minio"
)

var currentStorage Storage

type Storage interface {
	GetObjectToFile(objectKey, downloadedFileName string) error
	DeleteObject(objectKey string) error
	PutObject(objectKey string, reader io.Reader) error
	PutObjectFromFile(objectKey, filePath string) error
	IsExists(objectKey string) (bool, error)
	GetObjectUrl(objectKey string) string
}

func NewStorage(ossName string, config Config) error {
	if config.KeyId == "" || config.KeySecret == "" || config.Bucket == "" {
		return errors.New("configuration not correct")
	}
	var err error
	switch ossName {
	case AliYunConst:
		currentStorage, err = NewAliYun(config)
	case TencentConst:
		currentStorage, err = NewTencent(config)
	case QinNiuConst:
		currentStorage, err = NewQiNiu(config)
	case MinioConst:
		currentStorage, err = NewMinio(config)
	default:
		return errors.New("driver not exists")
	}
	if err != nil {
		currentStorage = nil
		return err
	}
	return nil
}

func GetStorage() (Storage, error) {
	if currentStorage != nil {
		return currentStorage, nil
	}
	return nil, errors.New("driver not exists")
}
