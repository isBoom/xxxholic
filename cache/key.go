package cache

import "fmt"

const (
	DailyRankKey = "rank:daily"
	WeekRankKey  = "rank:week"
	MonthRankKey = "rank:month"
)

var RankType = [...]string{DailyRankKey, WeekRankKey, MonthRankKey}

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
