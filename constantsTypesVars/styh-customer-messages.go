package sharedServices

//goland:noinspection All
const (
	// Categories
	CUST_MSG_NO_DATA_FOUND_FORMATTED = "Hmm, it looks like we couldn't find any information matching your request.\n " +
		"\u2022 Has your information been updated in your service provider.\n" +
		"\u2022 Try a broader search: Use more general keywords related to the SAAS provider."
	CUST_MSG_NEED_DATE_FORMATTED = "The question doesn't have a timeframe.\n" + "\u2022 Added a date.\n" +
		"\u2022 Does the date make sense in the context of the question. Example: A date in the future would not work for Sales questions.\n" +
		"\u2022 words like current, previous, today, yesterday\n" + "\u2022 last and the timeframe. ex: last week or last month"
)
