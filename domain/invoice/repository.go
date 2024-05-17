package invoice

type InvoiceRepository interface {
	CreateInvoice(invoice *Invoice) error
	GetInvoiceByID(id uint) (*Invoice, error)
	UpdateInvoice(invoice *Invoice) error
	DeleteInvoice(id uint) error
	ListInvoices() ([]*Invoice, error)
}
