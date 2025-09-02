package main

import (
	"github.com/OwnYoung/shortURL/controllers"
	"github.com/OwnYoung/shortURL/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDB()
	r := gin.Default()

	r.POST("/api/create", controllers.CreateShortLink)
	r.GET("/:code", controllers.RedirectShortLink)

	r.Run() // 默认在 localhost:8080 启动
}
