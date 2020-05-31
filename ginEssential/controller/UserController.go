package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "手机号没有11位",
		})
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "密码少于6位",
		})
	}
	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "名字长度不能为0",
		})
	}

	log.Println(name, telephone, password)

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func Login(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
