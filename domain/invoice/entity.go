package invoice

import "time"

// Invoice представляет собой сущность Invoice
type Invoice struct {
	ID             uint
	OrganizationID uint
	InvoiceNumber  string
	InvoiceDate    time.Time
	DueDate        time.Time
	Items          []InvoiceItem
	Total          float64
	PDFPath        string
}

// InvoiceItem представляет собой сущность InvoiceItem
type InvoiceItem struct {
	ID          uint
	InvoiceID   uint
	Description string
	Quantity    int
	UnitPrice   float64
	Total       float64
}
