package messages

type URL struct {
	LongURL   string `json:"longURL"`
	ShortPath string `json:"shortPath"`
}

type CreateURLCommand struct {
	BaseMessage

	LongURL string `json:"longURL"`
}

type ShortURLCreatedEvent struct {
	BaseMessage

	LongURL   string `json:"longURL"`
	ShortPath string `json:"shortPath"`
}

/**************************************************/
type DeleteURLCommand struct {
	BaseMessage

	ShortPath string `json:"shortPath"`
}

type ShortURLDeletedEvent struct {
	BaseMessage

	ShortPath string `json:"shortPath"`
}

/*************************************************/
type URLVisitedEvent struct {
	BaseMessage

	ShortPath string `json:"shortPath"`
	ETAG      string `json:"etag"`
	UserAgent string `json:"userAgent"`
}

/**************************************************/

type URLListQuery struct {
	BaseMessage

	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type URLListReadyEvent struct {
	BaseMessage

	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"Limit"`

	URLs []*URL `json:"urls"`
}

/****************************************************/

type LongURLQuery struct {
	BaseMessage

	ShortPath string `json:"shortPath"`
}

type LongURLReadyEvent struct {
	BaseMessage

	ShortPath string `json:"shortPath"`
	LongURL   string `json:"longURL"`
}
