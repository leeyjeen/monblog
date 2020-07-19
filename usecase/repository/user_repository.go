package repository

/* Use Cases Layer */
import (
	"github.com/leeyjeen/monblog/domain/model"
)

// UserRepository this is user repository interface
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

