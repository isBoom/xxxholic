package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
	"xxxholic/cache"
)

type Video struct {
	gorm.Model
	Title  string `json:"Title"`
	Info   string `json:"info"`
	Url    string `json:"url" form:"url"`
	Avatar string `json:"avatar"`
	UserId uint   `json:"userId"`
	VideoType string `json:"videoType"`
}

func (video *Video) AvatarUrl() string {
	client, _ := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKER"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 60)
	if (video.Avatar == "") || strings.Contains(signedGetURL, os.Getenv("OSS_UserInfoUrl")+"?Exp") {
		//signedGetURL = "https://xxxholic.oss-cn-hongkong.aliyuncs.com/upload/avatar/defaultAvatar.jpg"
		signedGetURL = ""
	}
	return signedGetURL
}

func (video *Video) VideoUrl() string {
	client, _ := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKER"))
	signedGetURL, _ := bucket.SignURL(video.Url, oss.HTTPGet, 3600)
	return signedGetURL
}
func (video *Video) AddView() {
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
func (video *Video) GetView() uint64 {
	count, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Uint64()
	return count
}
