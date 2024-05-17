package persistence

import (
	"github.com/Nikolay200669/go-auto-invoice/domain/cron_settings"
	"github.com/Nikolay200669/go-auto-invoice/domain/invoice"
	"github.com/Nikolay200669/go-auto-invoice/domain/organization"

	"github.com/jinzhu/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) (*GormRepository, error) {
	db.AutoMigrate(
		&invoice.Invoice{}, &invoice.InvoiceItem{},
		&organization.Organization{},
		&cron_settings.CronSettings{},
	)

	return &GormRepository{db: db}, nil
}

// Реализация методов репозиториев для Invoice, Organization и CronSettings

// Реализация методов репозитория для Invoice
func (r *GormRepository) ListInvoices() ([]*invoice.Invoice, error) {
	var invoices []*invoice.Invoice
	err := r.db.Find(&invoices).Error
	return invoices, err
}

func (r *GormRepository) CreateInvoice(invoice *invoice.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *GormRepository) GetInvoiceByID(id uint) (*invoice.Invoice, error) {
	var inv invoice.Invoice
	err := r.db.First(&inv, id).Error
	return &inv, err
}

func (r *GormRepository) UpdateInvoice(invoice *invoice.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *GormRepository) DeleteInvoice(id uint) error {
	return r.db.Delete(&invoice.Invoice{}, id).Error
}

// Реализация методов репозитория для Organization
func (r *GormRepository) CreateOrganization(org *organization.Organization) error {
	return r.db.Create(org).Error
}

func (r *GormRepository) GetOrganizationByID(id uint) (*organization.Organization, error) {
	var org organization.Organization
	err := r.db.First(&org, id).Error
	return &org, err
}

func (r *GormRepository) GetAllOrganizations() ([]organization.Organization, error) {
	var orgs []organization.Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r *GormRepository) UpdateOrganization(org *organization.Organization) error {
	return r.db.Save(org).Error
}

func (r *GormRepository) DeleteOrganization(id uint) error {
	return r.db.Delete(&organization.Organization{}, id).Error
}

// Реализация методов репозитория для CronSettings
func (r *GormRepository) CreateCronSettings(settings *cron_settings.CronSettings) error {
	return r.db.Create(settings).Error
}

func (r *GormRepository) GetCronSettingsByID(id uint) (*cron_settings.CronSettings, error) {
	var settings cron_settings.CronSettings
	err := r.db.First(&settings, id).Error
	return &settings, err
}

func (r *GormRepository) GetCronSettingsByOrganizationID(orgID uint) (*cron_settings.CronSettings, error) {
	var settings cron_settings.CronSettings
	err := r.db.Where("organization_id = ?", orgID).First(&settings).Error
	return &settings, err
}

func (r *GormRepository) UpdateCronSettings(settings *cron_settings.CronSettings) error {
	return r.db.Save(settings).Error
}

func (r *GormRepository) DeleteCronSettings(id uint) error {
	return r.db.Delete(&cron_settings.CronSettings{}, id).Error
}
