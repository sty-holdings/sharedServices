package sharedServices

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
	TBL_DAILY_PERFORMANCE  = "daily_performance"
	TBL_COUPLER_GOOGLE_ADS = "coupler_google_ads"
)

//goland:noinspection All
const (
	// Text strings
	INSERT_DAILY_PERFORMANCE = "INSERT INTO dkga.daily_performance " +
		"(campaign_id, campaign_type, campaign_name, date, clicks, impressions, ctr, cpc, spend, cpm, cost_per_conversion, conversion_rate, conversion_value) " +
		"VALUES (%v);\n"
	SELECT_ALL_FROM_TABLE = "SELECT * FROM %s.%s;\n"
)
