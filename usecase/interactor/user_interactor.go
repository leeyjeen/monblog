package interactor

/* Use Cases Layer */
import (
	"github.com/leeyjeen/monblog/domain/model"
	"github.com/leeyjeen/monblog/usecase/presenter"
	"github.com/leeyjeen/monblog/usecase/repository"
)

// UserInteractor this is user interactor interface
type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
}

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

// NewUserInteractor this is used to get new user interactor struct
func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (ui *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	users, err := ui.UserRepository.FindAll(u) // interactor uses repository
	if err != nil {
		return nil, err
	}
	return ui.UserPresenter.ResponseUsers(users), nil // interactor uses presenter
}
