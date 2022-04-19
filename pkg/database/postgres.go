package database

import (
	"ecommerce-boiler/internals/core/domain"
	"ecommerce-boiler/internals/core/ports"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataStore struct {
}

func NewDataBase() ports.IDatabase{
	return &dataStore{}
}

func (dt dataStore) ConnectDB(dsn string) *gorm.DB{
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("failed to connect database: %v", err)
	}
	fmt.Println("establishing connected to database")
	return db
}

func (dt dataStore) MigrateAll(db *domain.User) error {
	return dt.MigrateAll(&domain.User{})
}