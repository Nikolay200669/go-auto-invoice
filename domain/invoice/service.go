package invoice

type InvoiceService struct {
	repository InvoiceRepository
}

func NewInvoiceService(repo InvoiceRepository) *InvoiceService {
	return &InvoiceService{repository: repo}
}

func (s *InvoiceService) CreateInvoice(invoice *Invoice) error {
	return s.repository.CreateInvoice(invoice)
}

func (s *InvoiceService) GetInvoice(id uint) (*Invoice, error) {
	return s.repository.GetInvoiceByID(id)
}

func (s *InvoiceService) UpdateInvoice(invoice *Invoice) error {
	return s.repository.UpdateInvoice(invoice)
}

func (s *InvoiceService) DeleteInvoice(id uint) error {
	return s.repository.DeleteInvoice(id)
}

func (s *InvoiceService) ListInvoices() ([]*Invoice, error) {
	return s.repository.ListInvoices()
}
