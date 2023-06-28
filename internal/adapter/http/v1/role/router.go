package role

import (
	"github.com/gin-gonic/gin"
	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/internal/adapter"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
)

type roleRoutes struct {
	u application_role.IService
	l pwlogger.Interface
}

func NewRoleRoutes(handler *gin.RouterGroup, u application_role.IService, l pwlogger.Interface) {
	r := &roleRoutes{u, l}

	h := handler.Group("/role")
	{
		h.POST("/create", r.create)
	}
}

// @Summary     Create role
// @Description Create role
// @ID          CreateRole
// @Tags  	    Role
// @Accept      json
// @Produce     json
// @Param		body	body		dto.UserRequestDTO	true	"User Request"
// @Success		200		{object}	v1.Response
// @Failure		400		{object}	v1.Response
// @Failure		500		{object}	v1.Response
// @Router			/user/create [post]
func (r *roleRoutes) create(c *gin.Context) {
	// 開始追蹤
	// var tracer = otel.Tracer("")
	// ctx, span := tracer.Start(c.Request.Context(), "")
	// defer span.End()
	// // 參數驗證
	// // var req dto.UserRequestDTO
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	log.Error().Err(err).Msg("http - v1 - backend - CreateUser")
	// 	adapter.HandleErrorResponse(c, common_error.New(common_error.ErrorCodeParamMissing))

	// 	return
	// }
	// // 檢查檔案格式
	// if err := validator.New().Struct(req); err != nil {
	// 	log.Error().Err(err).Msg("http - v1 - backend - CreateUser")
	// 	adapter.HandleErrorResponse(c, common_error.New(common_error.ErrorCodeParamFormatInvalid))

	// 	return
	// }

	// // 設定追蹤屬性
	// if kv, err := operations.TransformToAttribute("request/", req); err == nil {
	// 	span.SetAttributes(kv...)
	// }
	// // 執行UseCase
	// var userReq domainDTO.UserInput
	// copier.Copy(&userReq, &req)
	// userReq.HospitalID = hospitalId
	// _, err = r.u.Create(ctx, &userReq)
	// if err != nil {
	// 	log.Error().Err(err).Msg("http - v1 - backend - CreateUser")
	// 	adapter.HandleErrorResponse(c, err)

	// 	return
	// }
	adapter.SuccessResponse(c, nil)
}
