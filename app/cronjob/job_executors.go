package cronjob

import (
	"github.com/robfig/cron"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func ExecCronJob(jobExecutor *JobExecutors) *cron.Cron {
	scheduler := setUpScheduler()
	log.Info().Msg("Exec Cron...")

	itemJobTime := viper.GetString("cron.item.sync")

	scheduler.AddFunc(itemJobTime, func() {
		log.Info().Msg("Exec job item sync...")
		jobExecutor.JobItems()
	})

	scheduler.Start()
	return scheduler
}
