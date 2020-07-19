package router

/* Frameworks and Drives Layer */
import (
	"github.com/gin-gonic/gin"
	"github.com/leeyjeen/monblog/interface/controller"
)

func NewRouter(c controller.AppController) *gin.Engine {
	r := gin.New()

	r.GET("/users", c.GetUsers)
	return r
}
