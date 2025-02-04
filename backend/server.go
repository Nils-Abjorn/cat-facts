package main

import (
	"github.com/Nils-Abjorn/cat-facts/controller"
	"github.com/Nils-Abjorn/cat-facts/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.FactsService       = service.New()
	videoController controller.FactsController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/facts", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.GET("/fact", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.GetRandomFact())
	})

	server.Run(":8080")
}
