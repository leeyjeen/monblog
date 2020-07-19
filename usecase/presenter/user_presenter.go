package presenter

/* Use Cases Layer */
import (
	"github.com/leeyjeen/monblog/domain/model"
)

// UserPresenter this is user presenter interface
type UserPresenter interface {
	ResponseUsers(users []*model.User) []*model.User
}
