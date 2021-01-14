package cache

import "fmt"

const (
	DailyRankKey = "rank:daily"
)

func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video:%d", id)
}
