package ports

import (
	"ecommerce-boiler/internals/core/domain"
	"gorm.io/gorm"
)

type IDatabase interface{
	ConnectDB (con string) *gorm.DB
	MigrateAll(db *domain.User) error
}
