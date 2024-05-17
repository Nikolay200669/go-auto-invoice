package organization

type OrganizationRepository interface {
	CreateOrganization(org *Organization) error
	GetOrganizationByID(id uint) (*Organization, error)
	GetAllOrganizations() ([]Organization, error)
	UpdateOrganization(org *Organization) error
	DeleteOrganization(id uint) error
}
