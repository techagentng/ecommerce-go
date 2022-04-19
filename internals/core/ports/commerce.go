package ports

import "ecommerce-boiler/internals/core/domain"
// ICompanyService defines the interface for company service
type ICommerceService interface {
	GetUserByID(id string) (*domain.User, error)
	GetAllUser(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateUser(company *domain.User) error
	UpdateUser(params common.GetByIDRequest, company common.UpdateCompanyRequest) (*domain.User, error)
	DeleteCompany(id string) error
}
