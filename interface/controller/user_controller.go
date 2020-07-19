package controller

/* Interface Adapter Layer */
import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leeyjeen/monblog/domain/model"
	"github.com/leeyjeen/monblog/usecase/interactor"
)

// UserController this is user controller interface
type UserController interface {
	GetUsers(c *gin.Context)
}

type userController struct {
	userInteractor interactor.UserInteractor
}

func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{ui}
}

func (uc *userController) GetUsers(c *gin.Context) {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, u)
}
