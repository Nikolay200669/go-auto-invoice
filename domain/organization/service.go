package organization

import "errors"

type OrganizationService struct {
	repository OrganizationRepository
}

func NewOrganizationService(repo OrganizationRepository) *OrganizationService {
	return &OrganizationService{repository: repo}
}

func (s *OrganizationService) CreateOrganization(org *Organization) error {
	if org.Name == "" {
		return errors.New("organization name is required")
	}

	return s.repository.CreateOrganization(org)
}

func (s *OrganizationService) GetOrganizationByID(id uint) (*Organization, error) {
	return s.repository.GetOrganizationByID(id)
}

func (s *OrganizationService) GetAllOrganizations() ([]Organization, error) {
	return s.repository.GetAllOrganizations()
}

func (s *OrganizationService) UpdateOrganization(org *Organization) error {
	if org.Name == "" {
		return errors.New("organization name is required")
	}

	return s.repository.UpdateOrganization(org)
}

func (s *OrganizationService) DeleteOrganization(id uint) error {
	return s.repository.DeleteOrganization(id)
}
