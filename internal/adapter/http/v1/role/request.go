package role

type CreatedRequest struct {
	Name        string   `json:"name" binding:"required" example:"admin"`
	Description string   `json:"description" binding:"required" example:"this is for admin role"`
	Permissions []string `json:"permissions" binding:"required" example:"read:all,write:all"`
}

type ListGotRequest struct {
	Limit      int      `json:"limit" form:"limit" binding:"required"`
	Offset     int      `json:"offset" form:"offset"`
	FilterName string   `json:"filterName" form:"filterName"`
	SortFields []string `json:"sortFields" form:"sortFields"`
	Dir        string   `json:"dir" form:"dir"`
}

type DetailGotRequest struct {
	ID string `json:"id" uri:"id" binding:"required"`
}

type DeletedRequest struct {
	ID string `json:"id" uri:"id" binding:"required"`
}

type UpdatedRequest struct {
	Name        string   `json:"name" binding:"required" example:"admin"`
	Description string   `json:"description" binding:"required" example:"this is for admin role"`
	Permissions []string `json:"permissions" binding:"required" example:"read:all,write:all"`
}
