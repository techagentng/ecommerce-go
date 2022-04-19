package ports

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ecommerce-boiler/internals/core/domain"
	"ecommerce-boiler/pkg/utils"
)

// IUserRepository defines the interface for user repositories
type IUserRepository interface {
	GetByID(id string) (*domain.User, error)
	GetBy(filter interface{}) ([]domain.User, error)
	Get(pagination *utils.Pagination) (*utils.Pagination, error)
	Persist(company *domain.User) error
	Delete(id string) error
	DeleteAll() error
	WithTx(tx *gorm.DB) IUserRepository
}

// IUserService defines the interface for user service
type IUserService interface {
	GetUserByID(id string) (interface{}, error)
	GetAllUser(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateUser(company *domain.User) error
	UpdateUser(params string) (*interface{}, error)
	DeleteUser(id string) error
}
// IUserHandler defines the interface for user handler
type IUserHandler interface {
	GetUserByID(c *gin.Context)
	GetAllUser(c *gin.Context)
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}