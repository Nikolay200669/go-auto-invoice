package repositories

import (
	"github.com/Nikolay200669/go-auto-invoice/domain/cron_settings"
	"github.com/Nikolay200669/go-auto-invoice/domain/invoice"
	"github.com/Nikolay200669/go-auto-invoice/domain/organization"
)

type Repository interface {
	invoice.InvoiceRepository
	organization.OrganizationRepository
	cron_settings.CronSettingsRepository
}
