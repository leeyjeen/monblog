package presenter

/* Interface Adapter Layer */
import (
	"github.com/leeyjeen/monblog/domain/model"
)

// UserPresenter this is user presenter interface
type UserPresenter interface {
	ResponseUsers(users []*model.User) []*model.User
}

type userPresenter struct {
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

// ResponseUsers handles user data before passing it to view
func (up *userPresenter) ResponseUsers(users []*model.User) []*model.User {
	for _, u := range users {
		u.Name = u.Name + "님"
	}
	return users
}
