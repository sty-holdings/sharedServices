package sharedServices

//goland:noinspection All
const (
	// Text strings
	INSERT_DAILY_PERFORMANCE = "INSERT INTO dkga.daily_performance " +
		"(campaign_id, campaign_type, campaign_name, date, clicks, impressions, ctr, cpc, spend, cpm, cost_per_conversion, conversion_rate, conversion_value) " +
		"VALUES (%v);\n"
)
