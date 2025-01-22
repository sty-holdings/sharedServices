package sharedServices

//goland:noinspection All
const (
	// Categories
	CUST_MSG_NO_DATA_FOUND_FORMATTED = `Hmm, it looks like we couldn't find any information matching your request.
		\u2022 Has your information been updated in your service provider.
		\u2022 Try a broader search: Use more general keywords related to the SAAS provider.`
	CUST_MSG_NEED_DATE_FORMATTED = `The question doesn't have a timeframe'.
		\u2022 Added a date.
        \u2022 Does the date make sense in the context of the question. Example: A date in the future would not work for Sales questions.
		\u2022 Added words like current, previous, today, yesterday
		\u2022 Added last and the timeframe. ex: last week or last month`
)
