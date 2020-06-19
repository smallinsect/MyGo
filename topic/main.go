package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context){

	})
	r.Run()
}
