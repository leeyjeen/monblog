package router

import (
	"github.com/gin-gonic/gin"
	"github.com/leeyjeen/monblog/interface/controller"
)

func NewRouter(c controller.AppController) *gin.Engine {
	r := gin.New()

	r.GET("/users", c.userController.GetUsers)
	return r
}
