package registry

/* registry plays a role in resolving dependencies using constructor injection */
import (
	"github.com/jinzhu/gorm"
	"github.com/leeyjeen/monblog/interface/controller"
)

// Registry this is registry interface
type Registry interface {
	NewAppController() controller.AppController
}

type registry struct {
	db *gorm.DB
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewUserController()
}
