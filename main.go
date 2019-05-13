package main

import (
	"github.com/gin-gonic/gin"
	"movie-show/commons"
	"movie-show/controllers"
	_ "movie-show/models"
)

func main() {
	router:=gin.Default()
	router.Use(commons.CheckDevice)
	apiRouter(router)
	router.Run(":8888")
}

func apiRouter(router *gin.Engine)  {
	api := router.Group("/api")
	{
		api.GET("/", controllers.ApiIndexGet)
		api.GET("/movie/:id", controllers.ApiDetailGet)
		api.GET("/play/:id/:num/:playType", controllers.ApiPlay)
		api.GET("/search", controllers.ApiSearch)
		api.GET("/test", controllers.ApiTest)
	}

}
