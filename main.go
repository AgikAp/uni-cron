package main

import (
	"admin_backend/app/cronjob"

	"github.com/gin-gonic/gin"
)

func main() {
	jobExecutor := cronjob.NewJobExecutors()
	scheduler := cronjob.ExecCronJob(jobExecutor)
	defer scheduler.Stop()

	c := gin.Default()
	c.Run(":8080")
}
