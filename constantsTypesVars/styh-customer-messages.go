package sharedServices

//goland:noinspection All
const (
	// Categories
	CUST_MSG_NO_DATA_FOUND_FORMATTED = "Hmm, it looks like we couldn't find any information matching your request.\nThe vendors you have connected to DaveKnows %s %s. If %s %s incorrect, " +
		"please enter a help desk ticket at https://www.daveknows.ai/Help.\n\nOther reasons:\n\tThere is no data for the question at this time." +
		"\n\tYour question is not related to the %s listed above.\n\tYour question is too specific."
	//
	CUST_MSG_NEED_DATE_FORMATTED = "The question doesn't have a timeframe, so I can't help you.\nPlease be more specific by adding a date or timeframe to your question.\n\nExample: in 2024, " +
		"10/29/2024, last week, last month, last quarter, etc."
)
