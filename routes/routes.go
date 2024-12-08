package routes

import (
	"github.com/gin-gonic/gin"
	"homework/infra"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	r := gin.Default()

	r.POST("register", ctx.Ctl.User.Registration)

	return r
}
