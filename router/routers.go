package router

import (
	"gin-admin/config"
	"gin-admin/controllers"
	"gin-admin/utils"
	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func InitRouter() *gin.Engine {

	cfg := config.NewConfig()
	route := gin.New()

	gin.SetMode(cfg.RunMode)
	route.Use(utils.GinRus(utils.Logger, time.RFC3339, false))
	route.Use(gin.Recovery())
	route.GET("/ping", controllers.Ping)
	route.GET("/user/:name", controllers.Authorized)

	route.GET("/metrics", gin.WrapH(promhttp.Handler()))
	authorized := route.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user: foo password:bar
		"manu": "123", // user:manu password: 123
	}))
	{
		authorized.POST("admin", controllers.Admin)
	}
	jobsRouter := route.Group("/job")
	{
		jobsRouter.POST("/add", controllers.AddJob)
		jobsRouter.GET("/list", controllers.ListJobs)
		jobsRouter.DELETE("/del/:jobID", controllers.RemoveJob)

	}

	jobrunner.Start()
	return route
}
