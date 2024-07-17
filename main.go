package main

import (
	"admin_backend/app/configs"
	"admin_backend/app/configs/logger"
	"admin_backend/app/cronjob"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.SetupLogger()
	configs.SetupConfig(".local")

	jobExecutor := cronjob.NewJobExecutors()
	scheduler := cronjob.ExecCronJob(jobExecutor)
	defer scheduler.Stop()

	c := gin.Default()
	c.Run(":8080")
}
