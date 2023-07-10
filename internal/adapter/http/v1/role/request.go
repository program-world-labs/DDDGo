package role

type CreatedRequest struct {
	Name        string   `json:"name" binding:"required" example:"admin"`
	Description string   `json:"description" binding:"required" example:"this is for admin role"`
	Permissions []string `json:"permissions" binding:"required" example:"read:all,write:all"`
}
