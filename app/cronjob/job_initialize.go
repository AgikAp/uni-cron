package cronjob

import (
	"time"

	"github.com/robfig/cron"
)

func setUpScheduler() *cron.Cron {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return cron.NewWithLocation(jakartaTime)
}
