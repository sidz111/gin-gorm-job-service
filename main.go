package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/job-service/config"
	"github.com/sidz111/job-service/controller"
	"github.com/sidz111/job-service/model"
	"github.com/sidz111/job-service/repository"
	"github.com/sidz111/job-service/service"
)

func main() {

	if err := config.ConnectDB(); err != nil {
		panic("failed to connect db")
	}
	router := gin.Default()

	config.DB.AutoMigrate(&model.Job{})
	repo := repository.NewJobRepository(config.DB)
	serv := service.NewJobService(repo)
	cont := controller.NewJobController(serv)
	job := router.Group("/jobs")
	{
		job.POST("/", cont.Create)
		job.GET("/:id", cont.GetById)
		job.GET("/", cont.GetAll)
		job.PUT("/", cont.Update)
		job.DELETE("/:id", cont.Delete)
	}
	router.Run(":8080")
}
