package main

import (
	"fmt"
	"gin-playground/src/controller"
	"gin-playground/src/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func PathParameter(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func GetVideo(ctx *gin.Context) {
	ctx.JSON(200, videoController.FindAll())
	videoController.FindAll()
}

func SaveVideos(ctx *gin.Context) {
	ctx.JSON(200, videoController.Save(ctx))
}

func main() {
	fmt.Println("Hello World")

	server := gin.Default()
	server.GET("/", HomePage)
	server.POST("/", PostHomePage)
	server.GET("/query", QueryStrings)
	server.GET("/path/:name/:age", PathParameter)

	server.GET("/videos", GetVideo)
	server.POST("/video", SaveVideos)

	server.Run()
}
