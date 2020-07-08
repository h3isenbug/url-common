package messages

type URLViewedEvent struct {
	Type      string `json:"type"`
	ShortPath string `json:"shortPath"`
	ETag      string `json:"etag"`
	UserAgent string `json:"userAgent"`
}
