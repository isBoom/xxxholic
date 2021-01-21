package task

import (
	"fmt"
	"github.com/robfig/cron"
	"reflect"
	"runtime"
	"time"
	"xxxholic/cache"
)

var Cron *cron.Cron

func Run(job func(name string) error, rankType string) {
	from := time.Now().UnixNano()
	err := job(rankType)
	to := time.Now().UnixNano()
	jobName := runtime.FuncForPC(reflect.ValueOf(job).Pointer()).Name()
	if err != nil {
		fmt.Printf("%s error: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	} else {
		fmt.Printf("%s success: %dms\n", jobName, (to-from)/int64(time.Millisecond))
	}
}
func CronJob() {
	if Cron == nil {
		Cron = cron.New()
	}
	Cron.AddFunc("@daily", func() {
		Run(RestarRank,cache.DailyRankKey)
	})
	Cron.AddFunc("@monthly", func() {
		Run(RestarRank,cache.MonthlyRankKey)
	})
	Cron.AddFunc("@weekly", func() {
		Run(RestarRank,cache.WeeklyRankKey)
	})
	//	Run(RestarRank,cache.WeekRankKey)
	//Run(RestarRank,cache.MonthRankKey)
	//Run(RestarRank,cache.DailyRankKey)
	Cron.Start()
	fmt.Println("Cron start...")
}
