package sharedServices

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

//goland:noinspection ALL
const (
	//
	PSQL_SSL_MODE_DISABLE     = "disable"
	PSQL_SSL_MODE_ALLOW       = "allow"
	PSQL_SSL_MODE_PREFER      = "prefer"
	PSQL_SSL_MODE_REQUIRED    = "require"
	PSQL_SSL_MODE_VERIFY      = "verify-ca"
	PSQL_SSL_MODE_VERIFY_FULL = "verify-full"
	//
	PSQL_CONN_STRING = "dbname=%s host=%s password=%s port=%d sslmode=%s sslrootcert=%s sslcert=%s sslkey=%s connect_timeout=%d user=%s"
	//
	SET_ROLE       = "SET ROLE %s;"
	TRUNCATE_TABLE = "TRUNCATE TABLE %s.%s;"
)

//goland:noinspection All
const (
	// Text strings
	INSERT_DAILY_PERFORMANCE = "INSERT INTO dkga.daily_performance " +
		"(campaign_id, campaign_type, campaign_name, date, clicks, impressions, ctr, cpc, spend, cpm, cost_per_conversion, conversion_rate, conversion_value) " +
		"VALUES (%v);\n"
	SELECT_ALL_FROM_TABLE = "SELECT * FROM %s.%s;\n"
	CHECK_STAT_ACTIVITY   = "SELECT * FROM pg_stat_activity WHERE datname = $1 and state = 'active';"
)

type PSQLConfig struct {
	GORM struct {
		UseGorm  bool `yaml:"use_gorm"`
		LoggerOn bool `yaml:"logger_on"`
	} `yaml:"gorm"`
	DBNames        []string     `json:"psql_db_names" yaml:"psql_db_names"`
	Debug          bool         `json:"psql_debug" yaml:"psql_debug"`
	Host           string       `json:"psql_host" yaml:"psql_host"`
	MaxConnections int          `json:"pgsql_max_connections" yaml:"pgsql_max_connections"`
	Password       string       `json:"psql_password" yaml:"psql_password"`
	Port           int          `json:"psql_port" yaml:"psql_port"`
	SSLMode        string       `json:"psql_ssl_mode" yaml:"psql_ssl_mode"`
	PSQLTLSInfo    jwts.TLSInfo `json:"psql_tls_info" yaml:"psql_tls_info"`
	Timeout        int          `json:"psql_timeout" yaml:"psql_timeout"`
	UserName       string       `json:"psql_user_name" yaml:"psql_user_name"`
}

type PSQLService struct {
	debugModeOn        bool
	ConnectionPoolPtrs map[string]*pgxpool.Pool
	GORMPoolPtrs       map[string]*gorm.DB
}

type DKCGACampaignPerformanceConsolidate struct {
	STYHInternalClientID        string  `db:"styh_client_id" json:"styh_client_id,omitempty" yaml:"styh_client_id"`
	AccountName                 string  `db:"account_name" json:"account_name,omitempty" yaml:"account_name"`
	CampaignName                string  `db:"campaign_name" json:"campaign_name,omitempty" yaml:"campaign_name"`
	AccountID                   string  `db:"account_id" json:"account_id,omitempty" yaml:"account_id"`
	CampaignID                  string  `db:"campaign_id" json:"campaign_id,omitempty" yaml:"campaign_id"`
	DkYear                      int     `db:"dk_year" json:"dk_year,omitempty" yaml:"dk_year"`
	DKQuarter                   int     `db:"dk_quarter" json:"dk_quarter,omitempty" yaml:"dk_quarter"`
	DKMonth                     int     `db:"dk_month" json:"dk_month,omitempty" yaml:"dk_month"`
	WeekOfDate                  string  `db:"week_of_date" json:"week_of_date,omitempty" yaml:"week_of_date"`
	ReportDate                  string  `db:"report_date" json:"report_date,omitempty" yaml:"report_date"`
	AmountSpend                 float64 `db:"amount_spend" json:"amount_spend,omitempty" yaml:"amount_spend"`
	CPC                         float64 `db:"cpc" json:"cpc,omitempty" yaml:"cpc"`
	CPM                         float64 `db:"cpm" json:"cpm,omitempty" yaml:"cpm"`
	CostPerConversion           float64 `db:"cost_per_conversion" json:"cost_per_conversion,omitempty" yaml:"cost_per_conversion"`
	Impressions                 float64 `db:"impressions" json:"impressions,omitempty" yaml:"impressions"`
	Clicks                      float64 `db:"clicks" json:"clicks,omitempty" yaml:"clicks"`
	CTR                         float64 `db:"ctr" json:"ctr,omitempty" yaml:"ctr"`
	Conversions                 float64 `db:"conversions" json:"conversions,omitempty" yaml:"conversions"`
	ViewThroughConversions      float64 `db:"view_through_conversions" json:"view_through_conversions,omitempty" yaml:"view_through_conversions"`
	TotalAmountSpend            float64 `db:"total_amount_spend" json:"total_amount_spend,omitempty" yaml:"total_amount_spend"`
	AvgCPC                      float64 `db:"avg_cpc" json:"avg_cpc,omitempty" yaml:"avg_cpc"`
	AvgCPM                      float64 `db:"avg_cpm" json:"avg_cpm,omitempty" yaml:"avg_cpm"`
	AvgCostPerConversion        float64 `db:"avg_cost_per_conversion" json:"avg_cost_per_conversion,omitempty" yaml:"avg_cost_per_conversion"`
	TotalImpressions            float64 `db:"total_impressions" json:"total_impressions,omitempty" yaml:"total_impressions"`
	TotalClicks                 float64 `db:"total_clicks" json:"total_clicks,omitempty" yaml:"total_clicks"`
	AvgCTR                      float64 `db:"avg_ctr" json:"avg_ctr,omitempty" yaml:"avg_ctr"`
	TotalConversions            float64 `db:"total_conversions" json:"total_conversions,omitempty" yaml:"total_conversions"`
	TotalViewThroughConversions float64 `db:"total_view_through_conversions" json:"total_view_through_conversions,omitempty" yaml:"total_view_through_conversions"`
}

// DK Tables go here
// DaveKnows tables used to provide information to users.

// AnalysisQuestionResults represents the dka.analysis_question_results table.
type AnalysisQuestionResults struct {
	AnalysisID              string    `db:"analysis_id" json:"analysis_id"`
	CreateTimestamp         time.Time `db:"create_timestamp" json:"create_timestamp"`
	AnalysisStatus          string    `db:"analysis_status" json:"analysis_status"`
	ElapseTimeSeconds       float64   `db:"elapse_time_seconds" json:"elapse_time_seconds"`
	TimePeriodDayValues     int64     `db:"time_period_day_values" json:"time_period_day_values"`
	TimePeriodMonthValues   int64     `db:"time_period_month_values" json:"time_period_month_values"`
	TimePeriodQuarterValues int64     `db:"time_period_quarter_values" json:"time_period_quarter_values"`
	TimePeriodWeekValues    int64     `db:"time_period_week_values" json:"time_period_week_values"`
	TimePeriodYearValues    int64     `db:"time_period_year_values" json:"time_period_year_values"`
	ComparisonQuestionFlag  bool      `db:"comparison_question_flag" json:"comparison_question_flag"`
	CurrentFlag             bool      `db:"current_flag" json:"current_flag"`
	DayFlag                 bool      `db:"day_flag" json:"day_flag"`
	DetailFlag              bool      `db:"detail_flag" json:"detail_flag"`
	LastFlag                bool      `db:"last_flag" json:"last_flag"`
	MonthFlag               bool      `db:"month_flag" json:"month_flag"`
	NextFlag                bool      `db:"next_flag" json:"next_flag"`
	PreviousFlag            bool      `db:"previous_flag" json:"previous_flag"`
	QuarterFlag             bool      `db:"quarter_flag" json:"quarter_flag"`
	SubtotalFlag            bool      `db:"subtotal_flag" json:"subtotal_flag"`
	TodayFlag               bool      `db:"today_flag" json:"today_flag"`
	TransactionFlag         bool      `db:"transaction_flag" json:"transaction_flag"`
	WeekFlag                bool      `db:"week_flag" json:"week_flag"`
	YearFlag                bool      `db:"year_flag" json:"year_flag"`
	Category                string    `db:"category" json:"category"`
	QuestionSubjects        string    `db:"question_subjects" json:"question_subjects"`
	GeneratedPrompt         string    `db:"generated_prompt" json:"generated_prompt"`
	UserQuestion            string    `db:"user_question" json:"user_question"`
}

// AnalysisTokenCounts represents the dka.analysis_token_counts table.
type AnalysisTokenCounts struct {
	AnalysisID          string `db:"analysis_id" json:"analysis_id"`
	TokenType           string `db:"token_type" json:"token_type"`
	CandidateTokenCount int64  `db:"candidate_token_count" json:"candidate_token_count"`
	PromptTokenCount    int64  `db:"prompt_token_count" json:"prompt_token_count"`
	TotalTokenCount     int64  `db:"total_token_count" json:"total_token_count"`
}

// GenerateAnswersResults represents the dka.generate_answers_results table.
type GenerateAnswersResults struct {
	AnalysisID        string    `db:"analysis_id" json:"analysis_id"`
	ElapseTimeSeconds float64   `db:"elapse_time_seconds" json:"elapse_time_seconds"`
	CreateTimestamp   time.Time `db:"create_timestamp" json:"create_timestamp"`
	Answer            string    `db:"answer" json:"answer"`
	AnswerStatus      string    `db:"answer_status" json:"answer_status"`
}

type Campaigns struct {
	styhInternalClientID   string    `db:"styh_client_id"` // UUID as string
	AccountAccountName     string    `db:"account_account_name"`
	CampaignCampaignName   string    `db:"campaign_campaign_name"`
	AccountAccountID       string    `db:"account_account_id"`
	CampaignCampaignID     string    `db:"campaign_campaign_id"`
	CampaignCampaignStatus string    `db:"campaign_campaign_status"`
	CampaignStartDate      time.Time `db:"campaign_start_date"`
	CampaignEndDate        time.Time `db:"campaign_end_date"`
	CampaignBudgetAmount   float64   `db:"campaign_budget_amount"`
}
type DKCGACampaignPerformanceDailyRow struct {
	styhInternalClientID   string  `db:"styh_client_id"`
	AccountName            string  `db:"account_name"`
	CampaignName           string  `db:"campaign_name"`
	ReportDate             string  `db:"report_date"`
	AccountID              string  `db:"account_id"`
	CampaignID             string  `db:"campaign_id"`
	AmountSpend            float64 `db:"amount_spend"`
	CPC                    float64 `db:"cpc"`
	CPM                    float64 `db:"cpm"`
	CostPerConversion      float64 `db:"cost_per_conversion"`
	Impressions            float64 `db:"impressions"`
	Clicks                 float64 `db:"clicks"`
	CTR                    float64 `db:"ctr"`
	Conversions            float64 `db:"conversions"`
	ViewThroughConversions float64 `db:"view_through_conversions"`
}

type DKCGACampaignPerformanceMonthlyRow struct {
	styhInternalClientID        string  `db:"styh_client_id"`
	AccountName                 string  `db:"account_name"`
	CampaignName                string  `db:"campaign_name"`
	dkYear                      int     `db:"dk_year"`
	dkMonth                     int     `db:"dk_month"`
	AccountID                   string  `db:"account_id"`
	CampaignID                  string  `db:"campaign_id"`
	TotalAmountSpend            float64 `db:"total_amount_spend"`
	AvgCPC                      float64 `db:"avg_cpc"`
	AvgCPM                      float64 `db:"avg_cpm"`
	AvgCostPerConversion        float64 `db:"avg_cost_per_conversion"`
	TotalImpressions            float64 `db:"total_impressions"`
	TotalClicks                 float64 `db:"total_clicks"`
	AvgCTR                      float64 `db:"avg_ctr"`
	TotalConversions            float64 `db:"total_conversions"`
	TotalViewThroughConversions float64 `db:"total_view_through_conversions"`
}

type DKCGACampaignPerformanceQuarterlyRow struct {
	styhInternalClientID        string  `db:"styh_client_id"`
	AccountName                 string  `db:"account_name"`
	CampaignName                string  `db:"campaign_name"`
	dkYear                      int     `db:"dk_year"`
	dkQuarter                   int     `db:"dk_quarter"`
	AccountID                   string  `db:"account_id"`
	CampaignID                  string  `db:"campaign_id"`
	TotalAmountSpend            float64 `db:"total_amount_spend"`
	AvgCPC                      float64 `db:"avg_cpc"`
	AvgCPM                      float64 `db:"avg_cpm"`
	AvgCostPerConversion        float64 `db:"avg_cost_per_conversion"`
	TotalImpressions            float64 `db:"total_impressions"`
	TotalClicks                 float64 `db:"total_clicks"`
	AvgCTR                      float64 `db:"avg_ctr"`
	TotalConversions            float64 `db:"total_conversions"`
	TotalViewThroughConversions float64 `db:"total_view_through_conversions"`
}

type DKCGACampaignPerformanceWeeklyRow struct {
	styhInternalClientID        string  `db:"styh_client_id"`
	AccountName                 string  `db:"account_name"`
	CampaignName                string  `db:"campaign_name"`
	WeekOfDate                  string  `db:"week_of_date"`
	AccountID                   string  `db:"account_id"`
	CampaignID                  string  `db:"campaign_id"`
	TotalAmountSpend            float64 `db:"total_amount_spend"`
	AvgCPC                      float64 `db:"avg_cpc"`
	AvgCPM                      float64 `db:"avg_cpm"`
	AvgCostPerConversion        float64 `db:"avg_cost_per_conversion"`
	TotalImpressions            float64 `db:"total_impressions"`
	TotalClicks                 float64 `db:"total_clicks"`
	AvgCTR                      float64 `db:"avg_ctr"`
	TotalConversions            float64 `db:"total_conversions"`
	TotalViewThroughConversions float64 `db:"total_view_through_conversions"`
}

type DKCGACampaignPerformanceYearlyRow struct {
	styhInternalClientID        string  `db:"styh_client_id"`
	AccountName                 string  `db:"account_name"`
	CampaignName                string  `db:"campaign_name"`
	dkYear                      int     `db:"dk_year"`
	AccountID                   string  `db:"account_id"`
	CampaignID                  string  `db:"campaign_id"`
	TotalAmountSpend            float64 `db:"total_amount_spend"`
	AvgCPC                      float64 `db:"avg_cpc"`
	AvgCPM                      float64 `db:"avg_cpm"`
	AvgCostPerConversion        float64 `db:"avg_cost_per_conversion"`
	TotalImpressions            float64 `db:"total_impressions"`
	TotalClicks                 float64 `db:"total_clicks"`
	AvgCTR                      float64 `db:"avg_ctr"`
	TotalConversions            float64 `db:"total_conversions"`
	TotalViewThroughConversions float64 `db:"total_view_through_conversions"`
}

// Coupler Tables go here
// Structs to support the Coupler.io create tables

type CouplerGoogleAdsRow struct {
	AccountName                                           string
	ActionItems                                           string
	AdGroup                                               string
	AdGroupAdAssetAutomationSettings                      string
	AdStrength                                            string
	Labels                                                string
	PolicySummaryApprovalStatus                           string
	PolicySummaryPolicyTopicEntries                       string
	PolicySummaryReviewStatus                             string
	PrimaryStatus                                         string
	PrimaryStatusReasons                                  string
	ResourceName                                          string
	Status                                                string
	AdGroupID                                             int64
	AddedByGoogleAds                                      bool
	AppAdDescriptions                                     string
	AppAdHeadlines                                        string
	AppAdHtml5MediaBundles                                string
	AppAdImages                                           string
	AppAdMandatoryAdText                                  string
	AppAdYoutubeVideos                                    string
	AppEngagementAdDescriptions                           string
	AppEngagementAdHeadlines                              string
	AppEngagementAdImages                                 string
	AppEngagementAdVideos                                 string
	AppPreRegistrationAdDescriptions                      string
	AppPreRegistrationAdHeadlines                         string
	AppPreRegistrationAdImages                            string
	AppPreRegistrationAdYoutubeVideos                     string
	CallAdBusinessName                                    string
	CallAdCallTracked                                     string
	CallAdConversionAction                                string
	CallAdConversionReportingState                        string
	CallAdCountryCode                                     string
	CallAdDescription1                                    string
	CallAdDescription2                                    string
	CallAdDisableCallConversion                           string
	CallAdHeadline1                                       string
	CallAdHeadline2                                       string
	CallAdPath1                                           string
	CallAdPath2                                           string
	CallAdPhoneNumber                                     string
	CallAdPhoneNumberVerificationURL                      string
	DemandGenCarouselAdBusinessName                       string
	DemandGenCarouselAdCallToActionText                   string
	DemandGenCarouselAdCarouselCards                      string
	DemandGenCarouselAdDescription                        string
	DemandGenCarouselAdHeadline                           string
	DemandGenCarouselAdLogoImage                          string
	DemandGenMultiAssetAdBusinessName                     string
	DemandGenMultiAssetAdCallToActionText                 string
	DemandGenMultiAssetAdDescriptions                     string
	DemandGenMultiAssetAdHeadlines                        string
	DemandGenMultiAssetAdLeadFormOnly                     string
	DemandGenMultiAssetAdLogoImages                       string
	DemandGenMultiAssetAdMarketingImages                  string
	DemandGenMultiAssetAdPortraitMarketingImages          string
	DemandGenMultiAssetAdSquareMarketingImages            string
	DemandGenProductAdBreadcrumb1                         string
	DemandGenProductAdBreadcrumb2                         string
	DemandGenProductAdBusinessName                        string
	DemandGenProductAdCallToAction                        string
	DemandGenProductAdDescription                         string
	DemandGenProductAdHeadline                            string
	DemandGenProductAdLogoImage                           string
	DemandGenVideoResponsiveAdBreadcrumb1                 string
	DemandGenVideoResponsiveAdBreadcrumb2                 string
	DemandGenVideoResponsiveAdBusinessName                string
	DemandGenVideoResponsiveAdCallToActions               string
	DemandGenVideoResponsiveAdDescriptions                string
	DemandGenVideoResponsiveAdHeadlines                   string
	DemandGenVideoResponsiveAdLogoImages                  string
	DemandGenVideoResponsiveAdLongHeadlines               string
	DemandGenVideoResponsiveAdVideos                      string
	DevicePreference                                      string
	DisplayUploadAdDisplayUploadProductType               string
	DisplayUploadAdMediaBundle                            string
	DisplayURL                                            string
	ExpandedDynamicSearchAdDescription                    string
	ExpandedDynamicSearchAdDescription2                   string
	ExpandedTextAdDescription                             string
	ExpandedTextAdDescription2                            string
	ExpandedTextAdHeadlinePart1                           string
	ExpandedTextAdHeadlinePart2                           string
	ExpandedTextAdHeadlinePart3                           string
	ExpandedTextAdPath1                                   string
	ExpandedTextAdPath2                                   string
	FinalAppURLs                                          string
	FinalMobileURLs                                       string
	FinalURLSuffix                                        string
	FinalURLs                                             string
	HotelAd                                               string
	ID                                                    int64
	ImageAdImageAssetAsset                                string
	ImageAdImageURL                                       string
	ImageAdMimeType                                       string
	ImageAdName                                           string
	ImageAdPixelHeight                                    int64
	ImageAdPixelWidth                                     int64
	ImageAdPreviewImageURL                                string
	ImageAdPreviewPixelHeight                             int64
	ImageAdPreviewPixelWidth                              int64
	LegacyAppInstallAd                                    string
	LegacyResponsiveDisplayAdAccentColor                  string
	LegacyResponsiveDisplayAdAllowFlexibleColor           bool
	LegacyResponsiveDisplayAdBusinessName                 string
	LegacyResponsiveDisplayAdCallToActionText             string
	LegacyResponsiveDisplayAdDescription                  string
	LegacyResponsiveDisplayAdFormatSetting                string
	LegacyResponsiveDisplayAdLogoImage                    string
	LegacyResponsiveDisplayAdLongHeadline                 string
	LegacyResponsiveDisplayAdMainColor                    string
	LegacyResponsiveDisplayAdMarketingImage               string
	LegacyResponsiveDisplayAdPricePrefix                  string
	LegacyResponsiveDisplayAdPromoText                    string
	LegacyResponsiveDisplayAdShortHeadline                string
	LegacyResponsiveDisplayAdSquareLogoImage              string
	LegacyResponsiveDisplayAdSquareMarketingImage         string
	LocalAdCallToActions                                  string
	LocalAdDescriptions                                   string
	LocalAdHeadlines                                      string
	LocalAdLogoImages                                     string
	LocalAdMarketingImages                                string
	LocalAdPath1                                          string
	LocalAdPath2                                          string
	LocalAdVideos                                         string
	Name                                                  string
	ResponsiveDisplayAdAccentColor                        string
	ResponsiveDisplayAdAllowFlexibleColor                 string
	ResponsiveDisplayAdBusinessName                       string
	ResponsiveDisplayAdCallToActionText                   string
	ResponsiveDisplayAdControlSpecEnableAssetEnhancements string
	ResponsiveDisplayAdControlSpecEnableAutogenVideo      string
	ResponsiveDisplayAdDescriptions                       string
	ResponsiveDisplayAdFormatSetting                      string
	ResponsiveDisplayAdHeadlines                          string
	ResponsiveDisplayAdLogoImages                         string
	ResponsiveDisplayAdLongHeadline                       string
	ResponsiveDisplayAdMainColor                          string
	ResponsiveDisplayAdMarketingImages                    string
	ResponsiveDisplayAdPricePrefix                        string
	ResponsiveDisplayAdPromoText                          string
	ResponsiveDisplayAdSquareLogoImages                   string
	ResponsiveDisplayAdSquareMarketingImages              string
	ResponsiveDisplayAdYoutubeVideos                      string
	ResponsiveSearchAdDescriptions                        string
	ResponsiveSearchAdHeadlines                           string
	ResponsiveSearchAdPath1                               string
	ResponsiveSearchAdPath2                               string
	ShoppingComparisonListingAdHeadline                   string
	ShoppingProductAd                                     string
	ShoppingSmartAd                                       string
	SmartCampaignAdDescriptions                           string
	SmartCampaignAdHeadlines                              string
	SystemManagedResourceSource                           string
	TextAdDescription1                                    string
	TextAdDescription2                                    string
	TextAdHeadline                                        string
	TrackingURLTemplate                                   string
	TravelAd                                              string
	Type                                                  string
	URLCollections                                        string
	URLCustomParameters                                   string
	VideoAdBumperActionButtonLabel                        string
	VideoAdBumperActionHeadline                           string
	VideoAdBumperCompanionBannerAsset                     string
	VideoAdInFeedDescription1                             string
	VideoAdInFeedDescription2                             string
	VideoAdInFeedHeadline                                 string
	VideoAdInFeedThumbnail                                string
	VideoAdInStreamActionButtonLabel                      string
	VideoAdInStreamActionHeadline                         string
	VideoAdInStreamCompanionBannerAsset                   string
	VideoAdNonSkippableActionButtonLabel                  string
	VideoAdNonSkippableActionHeadline                     string
	VideoAdNonSkippableCompanionBannerAsset               string
	VideoAdOutStreamDescription                           string
	VideoAdOutStreamHeadline                              string
	VideoAdVideoAsset                                     string
	VideoResponsiveAdBreadcrumb1                          string
	VideoResponsiveAdBreadcrumb2                          string
	VideoResponsiveAdCallToActions                        string
	VideoResponsiveAdCompanionBanners                     string
	VideoResponsiveAdDescriptions                         string
	VideoResponsiveAdHeadlines                            string
	VideoResponsiveAdLongHeadlines                        string
	VideoResponsiveAdVideos                               string
	CampaignID                                            int64
	CustomerID                                            int64
}

// This is used for the Coupler.io service
type CouplerCampaignPerformanceRow struct {
	AccountAccountName                string    `db:"account__account_name"`
	AccountAccountID                  string    `db:"account__account_id"`
	ReportDate                        time.Time `db:"report__date"`
	CampaignCampaignName              string    `db:"campaign__campaign_name"`
	CampaignCampaignID                string    `db:"campaign__campaign_id"`
	CampaignCampaignStatus            string    `db:"campaign__campaign_status"`
	CampaignStartDate                 time.Time `db:"campaign__start_date"`
	CampaignEndDate                   time.Time `db:"campaign__end_date"`
	CampaignBudgetAmount              float64   `db:"campaign__budget_amount"`
	CostAmountSpend                   float64   `db:"cost__amount_spend"`
	CostCPC                           float64   `db:"cost__cpc"`
	CostCPM                           float64   `db:"cost__cpm"`
	CostCostPerConversion             float64   `db:"cost__cost_per_conversion"`
	PerformanceImpressions            float64   `db:"performance__impressions"`
	PerformanceClicks                 float64   `db:"performance__clicks"`
	ClicksCTR                         float64   `db:"clicks__ctr"`
	ConversionsConversions            float64   `db:"conversions__conversions"`
	ConversionsViewThroughConversions float64   `db:"conversions__view_through_conversions"`
}

type CouplerGoogleAdsCampaignsRow struct {
	AccountName                                             string
	AccessibleBiddingStrategy                               string
	AdServingOptimizationStatus                             string
	AdvertisingChannelSubType                               string
	AdvertisingChannelType                                  string
	AppCampaignSettingAppID                                 string
	AppCampaignSettingAppStore                              string
	AppCampaignSettingBiddingStrategyGoalType               string
	AssetAutomationSettings                                 string
	AudienceSettingUseAudienceGrouped                       string
	BaseCampaign                                            string
	BiddingStrategy                                         string
	BiddingStrategySystemStatus                             string
	BiddingStrategyType                                     string
	BrandGuidelinesEnabled                                  bool
	CampaignBudget                                          string
	CampaignGroup                                           string
	CommissionCommissionRateMicros                          string
	DemandGenCampaignSettingsUpgradedTargeting              string
	DynamicSearchAdsSettingDomainName                       string
	DynamicSearchAdsSettingFeeds                            string
	DynamicSearchAdsSettingLanguageCode                     string
	DynamicSearchAdsSettingUseSuppliedUrlsOnly              bool
	EndDate                                                 time.Time
	ExcludedParentAssetFieldTypes                           string
	ExcludedParentAssetSetTypes                             string
	ExperimentType                                          string
	FinalURLSuffix                                          string
	FixedCpmGoal                                            string
	FixedCpmTargetFrequencyInfoTargetCount                  string
	FixedCpmTargetFrequencyInfoTimeUnit                     string
	FrequencyCaps                                           string
	GeoTargetTypeSettingNegativeGeoTargetType               string
	GeoTargetTypeSettingPositiveGeoTargetType               string
	HotelPropertyAssetSet                                   string
	HotelSettingHotelCenterID                               string
	ID                                                      int64
	KeywordMatchType                                        string
	Labels                                                  string
	ListingType                                             string
	LocalCampaignSettingLocationSourceType                  string
	LocalServicesCampaignSettingsCategoryBids               string
	ManualCPA                                               string
	ManualCPCEnhancedCPCEnabled                             bool
	ManualCPM                                               string
	ManualCPV                                               string
	MaximizeConversionValueTargetROAS                       string
	MaximizeConversionsTargetCpaMicros                      int64
	Name                                                    string
	NetworkSettingsTargetContentNetwork                     bool
	NetworkSettingsTargetGoogleSearch                       bool
	NetworkSettingsTargetGoogleTVNetwork                    bool
	NetworkSettingsTargetPartnerSearchNetwork               bool
	NetworkSettingsTargetSearchNetwork                      bool
	NetworkSettingsTargetYoutube                            bool
	OptimizationGoalSettingOptimizationGoalTypes            string
	OptimizationScore                                       float64
	PaymentMode                                             string
	PercentCPCCpcBidCeilingMicros                           string
	PercentCPCEnhancedCPCEnabled                            string
	PerformanceMaxUpgradePerformanceMaxCampaign             string
	PerformanceMaxUpgradePreUpgradeCampaign                 string
	PerformanceMaxUpgradeStatus                             string
	PrimaryStatus                                           string
	PrimaryStatusReasons                                    string
	RealTimeBiddingSettingOptIn                             string
	ResourceName                                            string
	SelectiveOptimizationConversionActions                  string
	ServingStatus                                           string
	ShoppingSettingAdvertisingPartnerIDs                    string
	ShoppingSettingCampaignPriority                         string
	ShoppingSettingDisableProductFeed                       string
	ShoppingSettingEnableLocal                              string
	ShoppingSettingFeedLabel                                string
	ShoppingSettingMerchantID                               string
	ShoppingSettingUseVehicleInventory                      string
	StartDate                                               time.Time
	Status                                                  string
	TargetCPACpcBidCeilingMicros                            string
	TargetCPACpcBidFloorMicros                              string
	TargetCPATargetCpaMicros                                int64
	TargetCPMTargetFrequencyGoalTargetCount                 string
	TargetCPMTargetFrequencyGoalTimeUnit                    string
	TargetCPV                                               string
	TargetImpressionShareCpcBidCeilingMicros                int64
	TargetImpressionShareLocation                           string
	TargetImpressionShareLocationFractionMicros             int64
	TargetROASCpcBidCeilingMicros                           string
	TargetROASCpcBidFloorMicros                             string
	TargetROASTargetROAS                                    string
	TargetSpendCpcBidCeilingMicros                          int64
	TargetSpendTargetSpendMicros                            float64
	TargetingSettingTargetRestrictions                      string
	TrackingSettingTrackingURL                              string
	TrackingURLTemplate                                     string
	TravelCampaignSettingsTravelAccountID                   string
	URLCustomParameters                                     string
	URLExpansionOptOut                                      string
	VanityPharmaVanityPharmaDisplayURLMode                  string
	VanityPharmaVanityPharmaText                            string
	VideoBrandSafetySuitability                             string
	VideoCampaignSettingsVideoAdInventoryControlAllowInFee  string
	VideoCampaignSettingsVideoAdInventoryControlAllowInStr  string
	VideoCampaignSettingsVideoAdInventoryControlAllowShorts string
	AccessibleBiddingStrategyID                             int64
	BiddingStrategyID                                       int64
	CampaignBudgetID                                        int64
	CampaignGroupID                                         string
	CustomerID                                              int64
}

type CouplerLinkedInCompanyRow struct {
	ReportPage                   string
	ReportPageID                 string
	ReportDate                   time.Time
	PerformanceImpressions       int64
	PerformanceUniqueImpressions int64
	PerformanceEngagementRate    float64
	PerformanceClicks            int64
	EngagementLikes              int64
	EngagementReposts            int64
	EngagementComments           int64
}
