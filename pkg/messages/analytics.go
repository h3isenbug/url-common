package messages

type URLViewedEvent struct {
	ShortPath string `json:"shortPath"`
	ETag      string `json:"etag"`
	UserAgent string `json:"userAgent"`
}
