package event

type RoleCreatedEvent struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions []string `json:"permissions"`
}

type RoleDescriptionChangedEvent struct {
	Description string `json:"description"`
}

type RolePermissionUpdatedEvent struct {
	Permissions []string `json:"permissions"`
}
