package sharedServices

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	jwts "github.com/sty-holdings/sharedServices/v2025/jwtServices"
)

//goland:noinspection ALL
const (
	PSQL_SSL_MODE_DISABLE     = "disable"
	PSQL_SSL_MODE_ALLOW       = "allow"
	PSQL_SSL_MODE_PREFER      = "prefer"
	PSQL_SSL_MODE_REQUIRED    = "require"
	PSQL_SSL_MODE_VERIFY      = "verify-ca"
	PSQL_SSL_MODE_VERIFY_FULL = "verify-full"
	//
	PSQL_CONN_STRING = "dbname=%s host=%s pool_max_conns=%d password=%s port=%d sslmode=%s connect_timeout=%d user=%s"
	//
	SET_ROLE       = "SET ROLE %s;\n"
	TRUNCATE_TABLE = "TRUNCATE TABLE %s.%s;\n"
)

//goland:noinspection All
const (
	ROLE_COUPLER_GOOGLE_ADS_ACCESS = "coupler_google_ads_access"
	ROLE_ANSWER_ACCESS             = "answers_access"
)

//goland:noinspection All
const (
	DB_ANSWERS    = "answers"
	DB_COUPLER    = "coupler"
	DB_GOOGLE_ADS = "google_ads"
)

//goland:noinspection All
const (
	SCHEMA_DKA  = "dka"
	SCHEMA_DKC  = "dkc"
	SCHEMA_DKGA = "dkga"
)

//goland:noinspection All
const (
	TBL_DAILY_PERFORMANCE            = "daily_performance"
	TBL_COUPLER_GOOGLE_ADS           = "coupler_google_ads"
	TBL_COUPLER_GOOGLE_ADS_CAMPAIGNS = "coupler_google_ads_campaigns"
	TBL_COUPLER_LINKEDIN_COMPANY     = "coupler_linkedin_company"
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
	DBName         []string     `json:"psql_db_names" yaml:"psql_db_names"`
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

type PSQLConnectionConfig struct {
	DBName         string
	Debug          bool
	Host           string
	MaxConnections int
	Password       string
	Port           int
	SSLMode        string
	PSQLTLSInfo    jwts.TLSInfo
	Timeout        int
	UserName       string
}

type PSQLService struct {
	DebugOn            bool
	ConnectionPoolPtrs map[string]*pgxpool.Pool
}

type DailyPerformance struct {
	CampaignID        int
	CampaignType      string
	CampaignName      string
	Date              time.Time
	Clicks            int
	Impressions       int64
	CTR               float64
	CPC               float64
	Spend             float64
	CPM               float64
	CostPerConversion float64
	ConversionRate    float64
	ConversionValue   float64
	CreatedAt         time.Time
}

type CouplerGoogleAds struct {
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

type CouplerGoogleAdsCampaigns struct {
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

type CouplerLinkedInCompany struct {
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

var (
	CGAAccountName                                           string
	CGAActionItems                                           string
	CGAAdGroup                                               string
	CGAAdGroupAdAssetAutomationSettings                      string
	CGAAdStrength                                            string
	CGALabels                                                string
	CGAPolicySummaryApprovalStatus                           string
	CGAPolicySummaryPolicyTopicEntries                       string
	CGAPolicySummaryReviewStatus                             string
	CGAPrimaryStatus                                         string
	CGAPrimaryStatusReasons                                  string
	CGAResourceName                                          string
	CGAStatus                                                string
	CGAAdGroupID                                             int64
	CGAAddedByGoogleAds                                      bool
	CGAAppAdDescriptions                                     string
	CGAAppAdHeadlines                                        string
	CGAAppAdHtml5MediaBundles                                string
	CGAAppAdImages                                           string
	CGAAppAdMandatoryAdText                                  string
	CGAAppAdYoutubeVideos                                    string
	CGAAppEngagementAdDescriptions                           string
	CGAAppEngagementAdHeadlines                              string
	CGAAppEngagementAdImages                                 string
	CGAAppEngagementAdVideos                                 string
	CGAAppPreRegistrationAdDescriptions                      string
	CGAAppPreRegistrationAdHeadlines                         string
	CGAAppPreRegistrationAdImages                            string
	CGAAppPreRegistrationAdYoutubeVideos                     string
	CGACallAdBusinessName                                    string
	CGACallAdCallTracked                                     string
	CGACallAdConversionAction                                string
	CGACallAdConversionReportingState                        string
	CGACallAdCountryCode                                     string
	CGACallAdDescription1                                    string
	CGACallAdDescription2                                    string
	CGACallAdDisableCallConversion                           string
	CGACallAdHeadline1                                       string
	CGACallAdHeadline2                                       string
	CGACallAdPath1                                           string
	CGACallAdPath2                                           string
	CGACallAdPhoneNumber                                     string
	CGACallAdPhoneNumberVerificationURL                      string
	CGADemandGenCarouselAdBusinessName                       string
	CGADemandGenCarouselAdCallToActionText                   string
	CGADemandGenCarouselAdCarouselCards                      string
	CGADemandGenCarouselAdDescription                        string
	CGADemandGenCarouselAdHeadline                           string
	CGADemandGenCarouselAdLogoImage                          string
	CGADemandGenMultiAssetAdBusinessName                     string
	CGADemandGenMultiAssetAdCallToActionText                 string
	CGADemandGenMultiAssetAdDescriptions                     string
	CGADemandGenMultiAssetAdHeadlines                        string
	CGADemandGenMultiAssetAdLeadFormOnly                     string
	CGADemandGenMultiAssetAdLogoImages                       string
	CGADemandGenMultiAssetAdMarketingImages                  string
	CGADemandGenMultiAssetAdPortraitMarketingImages          string
	CGADemandGenMultiAssetAdSquareMarketingImages            string
	CGADemandGenProductAdBreadcrumb1                         string
	CGADemandGenProductAdBreadcrumb2                         string
	CGADemandGenProductAdBusinessName                        string
	CGADemandGenProductAdCallToAction                        string
	CGADemandGenProductAdDescription                         string
	CGADemandGenProductAdHeadline                            string
	CGADemandGenProductAdLogoImage                           string
	CGADemandGenVideoResponsiveAdBreadcrumb1                 string
	CGADemandGenVideoResponsiveAdBreadcrumb2                 string
	CGADemandGenVideoResponsiveAdBusinessName                string
	CGADemandGenVideoResponsiveAdCallToActions               string
	CGADemandGenVideoResponsiveAdDescriptions                string
	CGADemandGenVideoResponsiveAdHeadlines                   string
	CGADemandGenVideoResponsiveAdLogoImages                  string
	CGADemandGenVideoResponsiveAdLongHeadlines               string
	CGADemandGenVideoResponsiveAdVideos                      string
	CGADevicePreference                                      string
	CGADisplayUploadAdDisplayUploadProductType               string
	CGADisplayUploadAdMediaBundle                            string
	CGADisplayURL                                            string
	CGAExpandedDynamicSearchAdDescription                    string
	CGAExpandedDynamicSearchAdDescription2                   string
	CGAExpandedTextAdDescription                             string
	CGAExpandedTextAdDescription2                            string
	CGAExpandedTextAdHeadlinePart1                           string
	CGAExpandedTextAdHeadlinePart2                           string
	CGAExpandedTextAdHeadlinePart3                           string
	CGAExpandedTextAdPath1                                   string
	CGAExpandedTextAdPath2                                   string
	CGAFinalAppURLs                                          string
	CGAFinalMobileURLs                                       string
	CGAFinalURLSuffix                                        string
	CGAFinalURLs                                             string
	CGAHotelAd                                               string
	CGAID                                                    int64
	CGAImageAdImageAssetAsset                                string
	CGAImageAdImageURL                                       string
	CGAImageAdMimeType                                       string
	CGAImageAdName                                           string
	CGAImageAdPixelHeight                                    int64
	CGAImageAdPixelWidth                                     int64
	CGAImageAdPreviewImageURL                                string
	CGAImageAdPreviewPixelHeight                             int64
	CGAImageAdPreviewPixelWidth                              int64
	CGALegacyAppInstallAd                                    string
	CGALegacyResponsiveDisplayAdAccentColor                  string
	CGALegacyResponsiveDisplayAdAllowFlexibleColor           bool
	CGALegacyResponsiveDisplayAdBusinessName                 string
	CGALegacyResponsiveDisplayAdCallToActionText             string
	CGALegacyResponsiveDisplayAdDescription                  string
	CGALegacyResponsiveDisplayAdFormatSetting                string
	CGALegacyResponsiveDisplayAdLogoImage                    string
	CGALegacyResponsiveDisplayAdLongHeadline                 string
	CGALegacyResponsiveDisplayAdMainColor                    string
	CGALegacyResponsiveDisplayAdMarketingImage               string
	CGALegacyResponsiveDisplayAdPricePrefix                  string
	CGALegacyResponsiveDisplayAdPromoText                    string
	CGALegacyResponsiveDisplayAdShortHeadline                string
	CGALegacyResponsiveDisplayAdSquareLogoImage              string
	CGALegacyResponsiveDisplayAdSquareMarketingImage         string
	CGALocalAdCallToActions                                  string
	CGALocalAdDescriptions                                   string
	CGALocalAdHeadlines                                      string
	CGALocalAdLogoImages                                     string
	CGALocalAdMarketingImages                                string
	CGALocalAdPath1                                          string
	CGALocalAdPath2                                          string
	CGALocalAdVideos                                         string
	CGAName                                                  string
	CGAResponsiveDisplayAdAccentColor                        string
	CGAResponsiveDisplayAdAllowFlexibleColor                 string
	CGAResponsiveDisplayAdBusinessName                       string
	CGAResponsiveDisplayAdCallToActionText                   string
	CGAResponsiveDisplayAdControlSpecEnableAssetEnhancements string
	CGAResponsiveDisplayAdControlSpecEnableAutogenVideo      string
	CGAResponsiveDisplayAdDescriptions                       string
	CGAResponsiveDisplayAdFormatSetting                      string
	CGAResponsiveDisplayAdHeadlines                          string
	CGAResponsiveDisplayAdLogoImages                         string
	CGAResponsiveDisplayAdLongHeadline                       string
	CGAResponsiveDisplayAdMainColor                          string
	CGAResponsiveDisplayAdMarketingImages                    string
	CGAResponsiveDisplayAdPricePrefix                        string
	CGAResponsiveDisplayAdPromoText                          string
	CGAResponsiveDisplayAdSquareLogoImages                   string
	CGAResponsiveDisplayAdSquareMarketingImages              string
	CGAResponsiveDisplayAdYoutubeVideos                      string
	CGAResponsiveSearchAdDescriptions                        string
	CGAResponsiveSearchAdHeadlines                           string
	CGAResponsiveSearchAdPath1                               string
	CGAResponsiveSearchAdPath2                               string
	CGAShoppingComparisonListingAdHeadline                   string
	CGAShoppingProductAd                                     string
	CGAShoppingSmartAd                                       string
	CGASmartCampaignAdDescriptions                           string
	CGASmartCampaignAdHeadlines                              string
	CGASystemManagedResourceSource                           string
	CGATextAdDescription1                                    string
	CGATextAdDescription2                                    string
	CGATextAdHeadline                                        string
	CGATrackingURLTemplate                                   string
	CGATravelAd                                              string
	CGAType                                                  string
	CGAURLCollections                                        string
	CGAURLCustomParameters                                   string
	CGAVideoAdBumperActionButtonLabel                        string
	CGAVideoAdBumperActionHeadline                           string
	CGAVideoAdBumperCompanionBannerAsset                     string
	CGAVideoAdInFeedDescription1                             string
	CGAVideoAdInFeedDescription2                             string
	CGAVideoAdInFeedHeadline                                 string
	CGAVideoAdInFeedThumbnail                                string
	CGAVideoAdInStreamActionButtonLabel                      string
	CGAVideoAdInStreamActionHeadline                         string
	CGAVideoAdInStreamCompanionBannerAsset                   string
	CGAVideoAdNonSkippableActionButtonLabel                  string
	CGAVideoAdNonSkippableActionHeadline                     string
	CGAVideoAdNonSkippableCompanionBannerAsset               string
	CGAVideoAdOutStreamDescription                           string
	CGAVideoAdOutStreamHeadline                              string
	CGAVideoAdVideoAsset                                     string
	CGAVideoResponsiveAdBreadcrumb1                          string
	CGAVideoResponsiveAdBreadcrumb2                          string
	CGAVideoResponsiveAdCallToActions                        string
	CGAVideoResponsiveAdCompanionBanners                     string
	CGAVideoResponsiveAdDescriptions                         string
	CGAVideoResponsiveAdHeadlines                            string
	CGAVideoResponsiveAdLongHeadlines                        string
	CGAVideoResponsiveAdVideos                               string
	CGACampaignID                                            int64
	CGACustomerID                                            int64
)
