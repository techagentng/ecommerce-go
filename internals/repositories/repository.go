package repositories

import (
	"ecommerce-boiler/internals/core/domain"
	_ "ecommerce-boiler/internals/core/ports"
	"ecommerce-boiler/pkg/utils"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Get(pagination *utils.Pagination) (*utils.Pagination, error) {
	var payload []string
	if err := r.db.Scopes(utils.Paginate(payload, pagination, r.db)).Find(&payload).Error; err != nil {
		fmt.Println(err.Error(), "repo")
		return nil, err
	}
	pagination.Rows = payload
	return pagination, nil
}

func (r *Repository) GetByID(id string) (interface{}, error) {
	var payload []string
	if err := r.db.Where("id = ?", id).First(&payload).Error; err != nil {
		return nil, err
	}
	return &payload, nil
}

func (r *Repository) GetBy(filter interface{}) (interface{}, error) {
	var payload []string
	if err := r.db.Model(&domain.User{}).Find(&payload, filter).Error; err != nil {
		return nil, err
	}
	return payload, nil
}

func (r *Repository) Persist(payload *interface{}) error {
	if err := r.db.Create(&payload).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(payload *interface{}) error {
	if err := r.db.Save(payload).Error; err != nil {
		return err
	}
	return nil

}

func (r *Repository) Delete(id string) error {
	return nil
}

func (r *Repository) DeleteAll(entity string) error {
	if err := r.db.Exec(fmt.Sprintf("DELETE FROM %v", entity)).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) WithTx(tx *gorm.DB) *Repository {
	return NewRepository(tx)
}
