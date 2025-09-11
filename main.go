package main

import (
	"github.com/OwnYoung/shortURL/controllers"
	"github.com/OwnYoung/shortURL/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	storage.InitDB()
	r := gin.Default()
	// 启用 CORS
	// TODO 处理静态页面
	r.Use(cors.Default())
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	// 短链生成和跳转接口
	r.POST("/post", controllers.CreateShortLink)
	r.GET("/:shortCode", controllers.RedirectShortLink)

	//? 避免 浏览器自动请求 /favicon.ico，用于显示标签页的小图标。
	//? 如果不处理，控制台会报错 404 not found
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // 或 c.File("./static/favicon.ico")
	})
	//? 也可以不处理

	r.Run() // 默认在 localhost:8080 启动
}
