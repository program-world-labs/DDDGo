package role

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"

	"github.com/program-world-labs/DDDGo/internal/adapter"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	common_error "github.com/program-world-labs/DDDGo/internal/domain/errors"
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
// @Param		body	body		CreatedRequest	true	"Role Create Request"
// @Success		200		{object}	adapter.Response{data=Response}
// @Failure		400		{object}	adapter.Response
// @Failure		500		{object}	adapter.Response
// @Router			/role/create [post]
func (r *roleRoutes) create(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer("")
	ctx, span := tracer.Start(c.Request.Context(), "")
	// // 設定追蹤屬性
	// if kv, err := operations.TransformToAttribute("request/", req); err == nil {
	// 	span.SetAttributes(kv...)
	// }

	defer span.End()

	// 參數驗證
	var req CreatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Err(err).Msg("http - v1 - backend - CreateUser")
		adapter.HandleErrorResponse(c, common_error.New(common_error.ErrorCodeParamMissing))

		return
	}
	// // 檢查檔案格式
	validate := validator.New()
	err := validate.RegisterValidation("custom_permission", CustomPermission)

	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - backend - CreateUser")
		adapter.HandleErrorResponse(c, common_error.New(common_error.ErrorCodeParamFormatInvalid))

		return
	}

	err = validate.Struct(req)

	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - backend - CreateUser")
		adapter.HandleErrorResponse(c, common_error.New(common_error.ErrorCodeParamFormatInvalid))

		return
	}

	// // 執行UseCase
	var input application_role.CreatedInput
	err = copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - backend - CreateUser")
		adapter.HandleErrorResponse(c, err)

		return
	}

	roleEntity, err := r.u.CreateRole(ctx, &input)
	if err != nil {
		r.l.Error().Err(err).Msg("http - v1 - backend - CreateUser")
		adapter.HandleErrorResponse(c, err)

		return
	}

	res := NewResponse(roleEntity)
	adapter.SuccessResponse(c, res)
}
