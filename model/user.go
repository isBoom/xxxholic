package model

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

// User 用户模型
type User struct {
	gorm.Model
	Email       string `gorm:"unique;size:30;not null"`
	PasswordDigest string
	UserName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
	Signature string `gorm:"size:50"`
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

func (user *User) AvatarUrl() string {
	client, _ := oss.New(os.Getenv("OSS_Endpoint"), os.Getenv("OSS_AccessKeyId"), os.Getenv("OSS_AccessKeySecret"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKER"))
	signedGetURL, _ := bucket.SignURL(user.Avatar, oss.HTTPGet, 60)
	if (user.Avatar == "") || strings.Contains(signedGetURL, os.Getenv("OSS_UserInfoUrl")+"?Exp") {
		//signedGetURL = "https://xxxholic.oss-cn-hongkong.aliyuncs.com/upload/avatar/defaultAvatar.jpg"
		signedGetURL = ""
	}
	return signedGetURL
}
