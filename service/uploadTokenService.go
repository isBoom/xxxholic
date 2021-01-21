package service

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"xxxholic/serializer"
)

type UploadTokenService struct {
	FileName string `json:"fileName" form:"fileName"`
}

func getContentType(s string) string {
	switch filepath.Ext(s) {
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".jpeg":
		return "image/jpeg"
	case ".jpg":
		return "image/jpeg"
	case ".mp4":
		return "video/mp4"
	}
	return ""
}
func (s *UploadTokenService) Post(src string) serializer.Response {
	client, err := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKER"))
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	options := []oss.Option{
		oss.ContentType(getContentType(s.FileName)),
	}

	key := src + uuid.Must(uuid.NewRandom()).String() + "_" + s.FileName

	signedPutUrl, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	signedGetUrl, err := bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  5002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutUrl,
			"get": signedGetUrl,
		},
	}
}
