package registry

import (
	"github.com/leeyjeen/monblog/interface/controller"
	ip "github.com/leeyjeen/monblog/interface/presenter"
	ir "github.com/leeyjeen/monblog/interface/repository"
	"github.com/leeyjeen/monblog/usecase/interactor"
	up "github.com/leeyjeen/monblog/usecase/presenter"
	ur "github.com/leeyjeen/monblog/usecase/repository"
)

// NewUserController generates a controller with interactor
func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

// NewUserInteractor returns an interactor with a repository and presenter
func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

// NewUserRepository returns repository of interface passed with a database instance
// which fulfills usecase interface
func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
