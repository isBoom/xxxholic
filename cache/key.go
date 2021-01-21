package cache

import "fmt"

const (
	DailyRankKey = "rank:daily"
	WeeklyRankKey  = "rank:weekly"
	MonthlyRankKey = "rank:monthly"
)

var RankType = [...]string{DailyRankKey, WeeklyRankKey, MonthlyRankKey}

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%d", id)
}
func GetType(rankType string) string {
	if rankType == "" {
		return DailyRankKey
	} else {
		return "rank:" + rankType
	}
}
func GetRankName(rank, videoType string) string {
	if videoType == "" {
		return rank
	} else {
		return rank + ":" + videoType
	}
}
func GetCaptcha(email string) string{
	return email + ":captcha"
}
func GetCaptchaTime(email string) string{
	return email + ":captcha:time"
}
