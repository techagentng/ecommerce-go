package services

import (
	"ecommerce-boiler/internals/repositories"
	log "github.com/sirupsen/logrus"

	"ecommerce-boiler/internals/core/domain"
	"ecommerce-boiler/internals/core/ports"
	"ecommerce-boiler/pkg/utils"
)

type userService struct {
	UserRepository repositories.Repository
	logger            *log.Logger
}

// NewCompanyService function create a new instance for service
func NewCompanyService(cr repositories.Repository, l *log.Logger) ports.IUserService {
	return &userService{
		UserRepository: cr,
		logger:            l,
	}
}

func (c *userService) GetUserByID(id string) (interface{}, error) {
	user, err := c.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *userService) GetUser(filter interface{}) (interface{}, error) {
	users, err := c.UserRepository.GetBy(filter)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return users, nil
}

func (c *userService) GetAllUser(pagination *utils.Pagination) (*utils.Pagination, error) {
	users, err := c.UserRepository.Get(pagination)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return users, nil
}

func (c *userService) CreateUser(user *domain.User) error {
	return nil
}

func (c userService) DeleteUser(id string) error {
	return nil
}

func (c *userService) UpdateUser(params string) (*interface{}, error) {
	return nil, nil
}
