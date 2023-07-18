package event

type UserCreatedEvent struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	EMail    string `json:"email"`
}

type UserPasswordChangedEvent struct {
	Password string `json:"password"`
}

type UserEmailChangedEvent struct {
	EMail string `json:"email"`
}
