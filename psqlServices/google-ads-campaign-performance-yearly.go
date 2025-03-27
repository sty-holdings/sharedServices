package sharedServices

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type CampaignPerformanceYearly struct {
	StyhClientID                string   `gorm:"type:uuid;not null"`
	AccountName                 string   `gorm:"type:VARCHAR(255);not null"`
	CampaignName                string   `gorm:"type:VARCHAR(255);not null"`
	DkYear                      int      `gorm:"type:integer;not null"`
	AccountID                   *string  `gorm:"type:VARCHAR(255)"`     // Use pointer for nullable
	CampaignID                  *string  `gorm:"type:VARCHAR(255)"`     // Use pointer for nullable
	TotalAmountSpend            *float64 `gorm:"type:double precision"` // Use pointer for nullable
	AvgCpc                      *float64 `gorm:"type:double precision"` // Use pointer for nullable
	AvgCpm                      *float64 `gorm:"type:double precision"` // Use pointer for nullable
	AvgCostPerConversion        *float64 `gorm:"type:double precision"` // Use pointer for nullable
	TotalImpressions            *float64 `gorm:"type:double precision"` // Use pointer for nullable
	TotalClicks                 *float64 `gorm:"type:double precision"` // Use pointer for nullable
	Ctr                         *float64 `gorm:"type:double precision"` // Use pointer for nullable
	TotalConversions            *float64 `gorm:"type:double precision"` // Use pointer for nullable
	TotalViewThroughConversions *float64 `gorm:"type:double precision"` // Use pointer for nullable
}

// TableName overrides the table name used by CampaignPerformanceYearly
func (CampaignPerformanceYearly) TableName() string {
	return "google_ads.campaign_performance_yearly"
}

func (psqlServicePtr *PSQLService) GetGoogleAdsYearlyData(styhClientId string, year int) (googleData string) {

	var (
		tx      *gorm.DB
		results []CampaignPerformanceYearly
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ? and dk_year = ?", styhClientId, year).Find(&results)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}
