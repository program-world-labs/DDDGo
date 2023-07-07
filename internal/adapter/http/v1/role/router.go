package role

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"

	"github.com/program-world-labs/DDDGo/internal/adapter/http"
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
// @Param		body	body		CreatedRequest	true	"Role Create Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/role/create [post].
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
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, NewBindJSONError(err))

		return
	}

	// // 執行UseCase
	var input application_role.CreatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, NewCopyError(err))

		return
	}

	input.Permissions = strings.Join(req.Permissions, ",")

	roleEntity, err := r.u.CreateRole(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - CreateRole")
		http.HandleErrorResponse(c, NewUsecaseError(err))

		return
	}

	res := NewResponse(roleEntity)
	http.SuccessResponse(c, res)
}
