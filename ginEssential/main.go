package main

import (
	"github.com/gin-gonic/gin"
	"xuqiulin.com/mygo/ginEssential/controller"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.Run() // listen and serve on 0.0.0.0:8080
}
