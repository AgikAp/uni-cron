package main

import (
	"admin_backend/app/cronjob"
)

func main() {
	jobExecutor := cronjob.NewJobExecutors()
	scheduler := cronjob.ExecCronJob(jobExecutor)
	defer scheduler.Stop()
}
