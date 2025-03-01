package sharedServices

//goland:noinspection All
const (
	// Categories
	CUST_MSG_NO_DATA_FOUND_FORMATTED = "Hmm, it looks like we couldn't find any information matching your request." +
		"1) Has your information been updated in your service provider." +
		"2) Try a broader search: Use more general keywords related to the SAAS provider."
	CUST_MSG_NEED_DATE_FORMATTED = "The question doesn't have a timeframe." + "1) Added a date." +
		"2) Does the date make sense in the context of the question. Example: A date in the future would not work for Sales questions." +
		"3) You can use words like current, previous, today, yesterday, last and the time period. ex: last week or last month"
)
