package respository

/* Interface Adapter Layer */
import (
	"github.com/jinzhu/gorm"
	"github.com/leeyjeen/monblog/domain/model"
)

// UserRepository this is user repository interface
type UserRepository interface {
	FindAll(users []*model.User) ([]*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// FindAll this is an actual implementation of the interface in usecase/repository/user_repository.go
// and it will be injected as a method of handling database...
func (ur *userRepository) FindAll(users []*model.User) ([]*model.User, error) {
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
