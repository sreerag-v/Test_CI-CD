package routes

import (
	"Test/Test-Crud/handler"

	"github.com/gin-gonic/gin"
)

func Router(ctx *gin.Engine) {
	ctx.POST("create-user", handler.CreateUser)
	ctx.GET("get-user/:id", handler.GetUserById)
	ctx.PUT("update-user/:id", handler.UpdateUserByid)

}
