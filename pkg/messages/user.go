package messages

type UserRegisterCommand struct {
	BaseMessage

	Username string`json:"username"`
	EMail    string`json:"eMail"`
	Password string`json:"password"`
}

type UserRegisteredEvent struct {
	BaseMessage
}

type UserRegistrationFailedEvent struct {
	BaseMessage

	Reason string`json:"reason"`
}

/***************************************************/

type UserLoginQuery struct {
	BaseMessage

	Username string`json:"username"`
	EMail    string`json:"eMail"`
	Password string`json:"password"`
}

type UserLoggedInSuccessfullyEvent struct {
	BaseMessage
}

type UserLoginFailedEvent struct {
	BaseMessage

	Reason string`json:"reason"`
}

type UserTokensGeneratedEvent struct {
	BaseMessage

	AuthToken    string`json:"authToken"`
	RefreshToken string`json:"refreshToken"`
}

/***************************************************/

type RenewAuthTokenQuery struct {
	BaseMessage

	RefreshToken string `json:"refreshToken"`
}

type AuthTokenRenewedEvent struct {
	BaseMessage

	AuthToken string `json:"authToken"`
}

/***************************************************/

type UserLogoutCommand struct {
	BaseMessage

	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}
