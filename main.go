package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haierspi/graphql-gorm-demo/db"
	"github.com/haierspi/graphql-gorm-demo/graphql"
	"io"
	"os"
)

func main() {

	// todo 不引入一下，db 模块 init 未执行，待看
	db.GetDB()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("graphql", graphql.Handler)

	r.Run(":8080")

}
