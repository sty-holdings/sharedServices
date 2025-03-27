package sharedServices

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoogleAdsCampaign struct {
	StyhClientID   uuid.UUID `gorm:"type:UUID"`
	AccountName    string    `gorm:"type:VARCHAR(255)"`
	CampaignName   string    `gorm:"type:VARCHAR(255)"`
	AccountID      *string   `gorm:"type:VARCHAR(255)"`
	CampaignID     *string   `gorm:"type:VARCHAR(255)"`
	CampaignStatus string    `gorm:"type:VARCHAR(255)"`
	StartDate      time.Time `gorm:"type:date;default:CURRENT_DATE"`
	EndDate        time.Time `gorm:"type:date;default:CURRENT_DATE"`
	BudgetAmount   float64   `gorm:"type:double precision;default:0.0"`
}

func (GoogleAdsCampaign) TableName() string {
	return "google_ads.campaigns"
}

type CampaignPerformanceDaily struct {
	StyhClientID           uuid.UUID `gorm:"type:uuid"`
	AccountName            string    `gorm:"type:VARCHAR(255)"`
	CampaignName           string    `gorm:"type:VARCHAR(255)"`
	ReportDate             time.Time `gorm:"type:date;default:CURRENT_DATE"`
	AccountID              *string   `gorm:"type:VARCHAR(255)"`
	CampaignID             *string   `gorm:"type:VARCHAR(255)"`
	AmountSpend            float64   `gorm:"type:double precision;default:0.0"`
	Cpc                    float64   `gorm:"type:double precision;default:0.0"`
	Cpm                    float64   `gorm:"type:double precision;default:0.0"`
	CostPerConversion      float64   `gorm:"type:double precision;default:0.0"`
	Impressions            float64   `gorm:"type:double precision;default:0.0"`
	Clicks                 float64   `gorm:"type:double precision;default:0.0"`
	Ctr                    float64   `gorm:"type:double precision;default:0.0"`
	Conversions            float64   `gorm:"type:double precision;default:0.0"`
	ViewThroughConversions float64   `gorm:"type:double precision;default:0.0"`
}

// TableName overrides the table name used by CampaignPerformanceDaily
func (CampaignPerformanceDaily) TableName() string {
	return "google_ads.campaign_performance_daily"
}

type CampaignPerformanceMonthly struct {
	StyhClientID                *string `gorm:"type:UUID"`
	AccountName                 *string `gorm:"type:VARCHAR(255)"`
	CampaignName                *string `gorm:"type:VARCHAR(255)"`
	DkYear                      int     `gorm:"type:integer;default:0"`
	DkMonth                     int     `gorm:"type:integer;default:0"`
	AccountID                   *string `gorm:"type:VARCHAR(255)"`
	CampaignID                  *string `gorm:"type:VARCHAR(255)"`
	TotalAmountSpend            float64 `gorm:"type:double precision;default:0.0"`
	AvgCpc                      float64 `gorm:"type:double precision;default:0.0"`
	AvgCpm                      float64 `gorm:"type:double precision;default:0.0"`
	AvgCostPerConversion        float64 `gorm:"type:double precision;default:0.0"`
	TotalImpressions            float64 `gorm:"type:double precision;default:0.0"`
	TotalClicks                 float64 `gorm:"type:double precision;default:0.0"`
	Ctr                         float64 `gorm:"type:double precision;default:0.0"`
	TotalConversions            float64 `gorm:"type:double precision;default:0.0"`
	TotalViewThroughConversions float64 `gorm:"type:double precision;default:0.0"`
}

// TableName overrides the table name used by CampaignPerformanceMonthly
func (CampaignPerformanceMonthly) TableName() string {
	return "google_ads.campaign_performance_monthly"
}

type CampaignPerformanceQuarterly struct {
	StyhClientID                *string `gorm:"type:UUID"`
	AccountName                 *string `gorm:"type:VARCHAR(255)"`
	CampaignName                *string `gorm:"type:VARCHAR(255)"`
	DkYear                      int     `gorm:"type:integer;default:0"`
	DkQuarter                   int     `gorm:"type:integer;default:0"`
	AccountID                   *string `gorm:"type:VARCHAR(255)"`
	CampaignID                  *string `gorm:"type:VARCHAR(255)"`
	TotalAmountSpend            float64 `gorm:"type:double precision;default:0.0"`
	AvgCpc                      float64 `gorm:"type:double precision;default:0.0"`
	AvgCpm                      float64 `gorm:"type:double precision;default:0.0"`
	AvgCostPerConversion        float64 `gorm:"type:double precision;default:0.0"`
	TotalImpressions            float64 `gorm:"type:double precision;default:0.0"`
	TotalClicks                 float64 `gorm:"type:double precision;default:0.0"`
	Ctr                         float64 `gorm:"type:double precision;default:0.0"`
	TotalConversions            float64 `gorm:"type:double precision;default:0.0"`
	TotalViewThroughConversions float64 `gorm:"type:double precision;default:0.0"`
}

// TableName overrides the table name used by CampaignPerformanceQuarterly
func (CampaignPerformanceQuarterly) TableName() string {
	return "google_ads.campaign_performance_quarterly"
}

type GoogleAdsCampaignPerformanceWeekly struct {
	StyhClientID                *uuid.UUID `gorm:"type:UUID"`
	AccountName                 *string    `gorm:"type:VARCHAR(255)"`
	CampaignName                *string    `gorm:"type:VARCHAR(255)"`
	WeekOfDate                  time.Time  `gorm:"type:date;default:CURRENT_DATE"`
	AccountID                   *string    `gorm:"type:VARCHAR(255)"`
	CampaignID                  *string    `gorm:"type:VARCHAR(255)"`
	TotalAmountSpend            float64    `gorm:"type:double precision;default:0.0"`
	AvgCpc                      float64    `gorm:"type:double precision;default:0.0"`
	AvgCpm                      float64    `gorm:"type:double precision;default:0.0"`
	AvgCostPerConversion        float64    `gorm:"type:double precision;default:0.0"`
	TotalImpressions            float64    `gorm:"type:double precision;default:0.0"`
	TotalClicks                 float64    `gorm:"type:double precision;default:0.0"`
	AvgCtr                      float64    `gorm:"type:double precision;default:0.0"`
	TotalConversions            float64    `gorm:"type:double precision;default:0.0"`
	TotalViewThroughConversions float64    `gorm:"type:double precision;default:0.0"`
}

// TableName overrides the table name used by GoogleAdsCampaignPerformanceWeekly
func (GoogleAdsCampaignPerformanceWeekly) TableName() string {
	return "google_ads.campaign_performance_weekly"
}

type CampaignPerformanceYearly struct {
	StyhClientID                string   `gorm:"type:uuid;not null"`
	AccountName                 string   `gorm:"type:VARCHAR(255);not null"`
	CampaignName                string   `gorm:"type:VARCHAR(255);not null"`
	DkYear                      int      `gorm:"type:integer;not null"`
	AccountID                   *string  `gorm:"type:VARCHAR(255)"`
	CampaignID                  *string  `gorm:"type:VARCHAR(255)"`
	TotalAmountSpend            *float64 `gorm:"type:double precision"`
	AvgCpc                      *float64 `gorm:"type:double precision"`
	AvgCpm                      *float64 `gorm:"type:double precision"`
	AvgCostPerConversion        *float64 `gorm:"type:double precision"`
	TotalImpressions            *float64 `gorm:"type:double precision"`
	TotalClicks                 *float64 `gorm:"type:double precision"`
	Ctr                         *float64 `gorm:"type:double precision"`
	TotalConversions            *float64 `gorm:"type:double precision"`
	TotalViewThroughConversions *float64 `gorm:"type:double precision"`
}

// TableName overrides the table name used by CampaignPerformanceYearly
func (CampaignPerformanceYearly) TableName() string {
	return "google_ads.campaign_performance_yearly"
}

func (psqlServicePtr *PSQLService) GetGoogleAdsCampaignData(styhClientId string, year int) (googleData []GoogleAdsCampaign) {

	var (
		tx *gorm.DB
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ?", styhClientId, year).Find(&googleData)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}

func (psqlServicePtr *PSQLService) GetGoogleAdsYearData(styhClientId string, year int) (googleData []CampaignPerformanceYearly) {

	var (
		tx *gorm.DB
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ? and dk_year = ?", styhClientId, year).Find(&googleData)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}

func (psqlServicePtr *PSQLService) GetGoogleAdsQuarterData(styhClientId string, year int, quarter int) (googleData []CampaignPerformanceQuarterly) {

	var (
		tx *gorm.DB
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ? and dk_year = ? and dk_quarter = ?", styhClientId, year, quarter).Find(&googleData)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}

func (psqlServicePtr *PSQLService) GetGoogleAdsMonthData(styhClientId string, year int, month int) (googleData []CampaignPerformanceMonthly) {

	var (
		tx *gorm.DB
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ? and dk_year = ? and dk_month = ?", styhClientId, year, month).Find(&googleData)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}

func (psqlServicePtr *PSQLService) GetGoogleAdsWeekData(styhClientId string, startOfWeek string) (googleData []GoogleAdsCampaignPerformanceWeekly) {

	var (
		tx *gorm.DB
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ? and start_of_week = ?", styhClientId, startOfWeek).Find(&googleData)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}

func (psqlServicePtr *PSQLService) GetGoogleAdsDayData(styhClientId string, year int, month int, day int) (googleData []CampaignPerformanceDaily) {

	var (
		tx *gorm.DB
	)

	tx = psqlServicePtr.GORMPoolPtrs[DATABASE_ANSWERS].Where("styh_client_id = ? and year = ? and month = ? and day = ?", styhClientId, year, month, day).Find(&googleData)

	if psqlServicePtr.DebugOn {
		log.Printf("Rows: %d\n", tx.RowsAffected)
	}

	// Handle the error
	if tx.Error != nil {
		fmt.Printf("--> Error: %s", tx.Error.Error())
	}

	return
}
