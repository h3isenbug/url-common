package messages

type baseAnalyticsQuery struct {
	BaseMessage

	Period   string `json:"period"` //today, past-1-day, past-7-days, past-30-days
	ShortURL string `json:"shortURL"`
}

type URLOverallAnalyticsQuery struct {
	baseAnalyticsQuery
}

type URLPerUniqueUserAnalyticsQuery struct {
	baseAnalyticsQuery
}

type baseAnalyticsReadyEvent struct {
	BaseMessage

	Period   string `json:"period"` //today, past-1-day, past-7-days, past-30-days
	ShortURL string `json:"shortURL"`

	OverallVisits       uint64 `json:"overallVisits"`
	MobileVisits        uint64 `json:"mobileVisits"`
	DesktopVisits       uint64 `json:"desktopVisits"`
	OtherPlatformVisits uint64 `json:"otherPlatformVisits"`

	BrowserVisits map[string]uint64 `json:"browserVisits"`
}

type URLOverallAnalyticsReadyEvent struct {
	baseAnalyticsQuery
}

type URLPerUniqueUserAnalyticsReadyEvent struct {
	baseAnalyticsQuery
}
