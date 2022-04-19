package services

import (
	"ecommerce-boiler/internals/repositories"
	"errors"
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

func (c *userService) GetUserByID(id string) (*domain.User, error) {
	company, err := c.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (c *userService) GetUser(filter interface{}) ([]domain.User, error) {
	companies, err := c.UserRepository.GetBy(filter)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return companies, nil
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
	var entity []domain.User
	entity, err := c.UserRepository.GetBy(domain.User{ID: user.ID, Name: user.Name})
	if err != nil {
		return err
	}

	if len(entity) > 0 {
		return errors.New("already exist")
	}

	err = c.UserRepository.Persist(user)

	if err != nil {
		c.logger.Error(err)
		return err
	}

	return nil
}

func (c userService) DeleteUser(id string) error {
	err := c.UserRepository.Delete(id)
	if err != nil {
		c.logger.Error(err)
		return err
	}
	return nil
}

func (c *userService) UpdateUser(params string) (*domain.User, error) {
	user, err := c.UserRepository.GetByID(params)
	if err != nil {
		c.logger.Error(err)
		return nil, err
	}
	return user, nil
}
