package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
	"time"
	"xxxholic/cache"
)

type Video struct {
	gorm.Model
	Title     string `json:"Title"`
	Info      string `json:"info"`
	Url       string `json:"url" form:"url"`
	Avatar    string `json:"avatar"`
	UserId    uint   `json:"userId"`
	VideoType string `json:"videoType"`
	Status string `json:"status"`
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
	for _, value := range cache.RankType {
		cache.RedisClient.ZIncrBy(value, 1, strconv.Itoa(int(video.ID)))
		cache.RedisClient.ZIncrBy(cache.GetRankName(value, video.VideoType), 1, strconv.Itoa(int(video.ID)))
	}
}
func (video *Video) SaveHistory(userId uint){
	if userId != 0{
		cache.RedisClient.ZAdd(cache.GetHistoryName(userId),redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: strconv.Itoa(int(video.ID)),
		})
		//cache.RedisClient.ZIncrBy(cache.GetHistoryName(userId),float64(time.Now().Unix()),strconv.Itoa(int(video.ID)))
	}
}
func (video *Video) GetView() uint64 {
	count, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Uint64()
	return count
}
