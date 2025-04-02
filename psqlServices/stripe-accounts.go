package sharedServices

import (
	"time"
)

type DKPullsStripeAccounts struct {
	ID                   string    `gorm:"primaryKey;type:varchar(255)" json:"id"`
	ChargesEnabled       bool      `gorm:"type:boolean" json:"charges_enabled"`
	Created              time.Time `gorm:"type:date;default:CURRENT_DATE" json:"created"`
	Deleted              bool      `gorm:"type:boolean" json:"deleted"`
	DirectorsProvided    bool      `gorm:"type:boolean" json:"directors_provided"`
	EstimatedWorkerCount int64     `gorm:"type:bigint" json:"estimated_worker_count"`
	ExecutivesProvided   bool      `gorm:"type:boolean" json:"executives_provided"`
	PayoutsEnabled       bool      `gorm:"type:boolean" json:"payouts_enabled"`
	TaxIDProvided        bool      `gorm:"type:boolean" json:"tax_id_provided"`
	VatIDProvided        bool      `gorm:"type:boolean" json:"vat_id_provided"`
	Amount               int64     `gorm:"type:bigint" json:"amount"`
	Country              string    `gorm:"type:varchar(2)" json:"country"`
	DefaultCurrency      string    `gorm:"type:varchar(3)" json:"default_currency"`
	Currency             string    `gorm:"type:varchar(3)" json:"currency"`
	FiscalYearEnd        string    `gorm:"type:varchar(10)" json:"fiscal_year_end"`
	Name                 string    `gorm:"type:varchar(255)" json:"name"`
	Email                string    `gorm:"type:varchar(255)" json:"email"`
	Phone                string    `gorm:"type:varchar(255)" json:"phone"`
	Type                 string    `gorm:"type:varchar(255)" json:"type"`
	City                 string    `gorm:"type:varchar(255)" json:"city"`
	Line1                string    `gorm:"type:varchar(255)" json:"line1"`
	Line2                string    `gorm:"type:varchar(255)" json:"line2"`
	PostalCode           string    `gorm:"type:varchar(255)" json:"postal_code"`
	State                string    `gorm:"type:varchar(255)" json:"state"`
	URL                  string    `gorm:"type:varchar(255)" json:"url"`
	BusinessType         string    `gorm:"type:varchar(255)" json:"business_type"`
	Structure            string    `gorm:"type:varchar(255)" json:"structure"`
	SupportEmail         string    `gorm:"type:varchar(255)" json:"support_email"`
	SupportPhone         string    `gorm:"type:varchar(255)" json:"support_phone"`
	SupportURL           string    `gorm:"type:varchar(255)" json:"support_url"`
	ExportLicenseID      string    `gorm:"type:varchar(255)" json:"export_license_id"`
	ExportPurposeCode    string    `gorm:"type:varchar(255)" json:"export_purpose_code"`
}

// TableName overrides the table name used by DKPullsStripeAccounts
func (DKPullsStripeAccounts) TableName() string {
	return "dk_pulls.stripe_accounts"
}
