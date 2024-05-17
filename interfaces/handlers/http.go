package handlers

import (
	"context"

	"cdr.dev/slog"
	"github.com/Nikolay200669/go-auto-invoice/domain/cron_settings"
	"github.com/Nikolay200669/go-auto-invoice/domain/invoice"
	"github.com/Nikolay200669/go-auto-invoice/domain/organization"
	"github.com/Nikolay200669/go-auto-invoice/interfaces/repositories"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	invoiceService      *invoice.InvoiceService
	organizationService *organization.OrganizationService
	cronSettingsService *cron_settings.CronSettingsService
	logger              *slog.Logger
}

func NewHTTPServer(repo repositories.Repository, logger *slog.Logger) *HTTPServer {
	return &HTTPServer{
		invoiceService:      invoice.NewInvoiceService(repo),
		organizationService: organization.NewOrganizationService(repo),
		cronSettingsService: cron_settings.NewCronSettingsService(repo),
		logger:              logger,
	}
}

func (s *HTTPServer) Start() {
	router := gin.Default()

	router.GET("/invoices", s.handleInvoices)
	router.GET("/organizations", s.handleOrganizations)
	router.GET("/cron-settings", s.handleCronSettings)

	router.Run(":8080")
}

// Обработчики HTTP-запросов для Invoice, Organization и CronSettings
func (s *HTTPServer) handleInvoices(c *gin.Context) {
	s.logger.Info(context.Background(), "Handling invoice request")
	// Обработка запросов, связанных с Invoice
	items, err := s.invoiceService.ListInvoices()
	if err != nil {
		s.logger.Error(context.Background(), "Failed to create invoice", "error", err)
		c.JSON(500, gin.H{"error": "Failed to create invoice"})
		return
	}

	// Return successful response for GET
	c.JSON(200, gin.H{"message": items})
}

func (s *HTTPServer) handleOrganizations(c *gin.Context) {
	s.logger.Info(context.Background(), "Handling organization request")
	// Обработка запросов, связанных с Organization
}

func (s *HTTPServer) handleCronSettings(c *gin.Context) {
	s.logger.Info(context.Background(), "Handling cron settings request")
	// Обработка запросов, связанных с CronSettings
}
