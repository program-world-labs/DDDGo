package user

type CreatedRequest struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	EMail       string `json:"email"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	// RoleIDs     []string `json:"roleIds"`
	// GroupID string `json:"groupId"`
}

type ListGotRequest struct {
	Limit      int      `json:"limit" form:"limit" binding:"required"`
	Offset     int      `json:"offset" form:"offset"`
	FilterName string   `json:"filterName" form:"filterName"`
	SortFields []string `json:"sortFields" form:"sortFields"`
	Dir        string   `json:"dir" form:"dir"`
}

type DetailRequest struct {
	ID string `json:"id" uri:"id" binding:"required"`
}

type UpdateRequest struct {
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
}

type DeleteRequest struct {
	ID string `json:"id" uri:"id" binding:"required"`
}
