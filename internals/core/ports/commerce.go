package ports

import "ecommerce-boilerplate/internals/core/domain"
// ICompanyService defines the interface for company service
type ICommerceService interface {
	GetCompanyByID(id string) (*domain.User, error)
	GetAllCompany(pagination *utils.Pagination) (*utils.Pagination, error)
	CreateCompany(company *domain.Company) error
	UpdateCompany(params common.GetByIDRequest, company common.UpdateCompanyRequest) (*domain.Company, error)
	DeleteCompany(id string) error
}
