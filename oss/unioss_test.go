//============================================================
// 作者: 杨大雷
// 日期: 2023/12/02 8:32 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package oss

import (
	"fmt"
	"testing"
)

func TestOSS(t *testing.T) {
	err := NewStorage(MinioConst, Config{
		KeyId:     "Xt5eqli0e69gKRwS",
		KeySecret: "nScOjwrRMSqvYDA8sEV9Sde9tF7iBL3L",
		Endpoint:  "127.0.0.1:9000",
		Bucket:    "minio-go",
		UseSSL:    false,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, err := GetStorage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = s.PutObjectFromFile("testMinio", "/Users/maktub/Desktop/background/IMG_4530.PNG")
	if err != nil {
		fmt.Println(err.Error())
	}
	url := s.GetObjectUrl("testMinio")
	fmt.Println(url)
}
