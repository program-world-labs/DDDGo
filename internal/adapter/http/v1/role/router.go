package role

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"

	"github.com/program-world-labs/DDDGo/internal/adapter/http"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
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
		h.GET("/list", r.list)
		h.GET("/detail/:id", r.detail)
		h.PUT("/update/:id", r.update)
		// h.DELETE("/delete/:id", r.delete)
		// h.PUT("/assign-role/:id", r.assignRole)
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
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleBindJSON, err))

		return
	}

	// // 執行UseCase
	var input application_role.CreatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleCopyToInput, err))

		return
	}

	input.Permissions = strings.Join(req.Permissions, ",")

	roleEntity, err := r.u.CreateRole(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - CreateRole")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleUsecase, err))

		return
	}

	http.SuccessResponse(c, NewResponse(roleEntity))
}

// @Summary     List role
// @Description List role
// @ID          ListRole
// @Tags  	    Role
// @Accept      json
// @Produce     json
// @Param		limit	query	int		false	"Limit"
// @Param		offset	query	int		false	"Offset"
// @Param		filterName	query	string		false	"FilterName"
// @Param		sortFields	query	string		false	"SortFields"
// @Param		dir	query	string		false	"Dir"
// @Success		200		{object}	http.Response{data=ResponseList}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/role/list [get].
func (r *roleRoutes) list(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer("")
	ctx, span := tracer.Start(c.Request.Context(), "")
	// // 設定追蹤屬性
	// if kv, err := operations.TransformToAttribute("request/", req); err == nil {
	// 	span.SetAttributes(kv...)
	// }

	defer span.End()

	// 參數驗證
	var req ListGotInput
	if err := c.ShouldBindQuery(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindQuery")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleBindQuery, err))

		return
	}

	var input application_role.ListGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleCopyToInput, err))

		return
	}

	// // 執行UseCase
	roleEntities, err := r.u.GetRoleList(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - ListRole")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleUsecase, err))

		return
	}

	http.SuccessResponse(c, NewResponseList(roleEntities))
}

// @Summary     Detail role
// @Description Detail role
// @ID          DetailRole
// @Tags  	    Role
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/role/detail/{id} [get].
func (r *roleRoutes) detail(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer("")
	ctx, span := tracer.Start(c.Request.Context(), "")
	// // 設定追蹤屬性
	// if kv, err := operations.TransformToAttribute("request/", req); err == nil {
	// 	span.SetAttributes(kv...)
	// }

	defer span.End()

	// 參數驗證
	var req DetailGotInput
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleBindQuery, err))

		return
	}

	var input application_role.DetailGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleCopyToInput, err))

		return
	}

	// 執行UseCase
	roleEntity, err := r.u.GetRoleDetail(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DetailRole")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleUsecase, err))

		return
	}

	http.SuccessResponse(c, NewResponse(roleEntity))
}

// @Summary     Update role
// @Description Update role
// @ID          UpdateRole
// @Tags  	    Role
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Param		body	body		UpdatedRequest	true	"Role Update Request"
// @Success		200		{object}	http.Response
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/role/update/{id} [put].
func (r *roleRoutes) update(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer("")
	ctx, span := tracer.Start(c.Request.Context(), "")
	// // 設定追蹤屬性
	// if kv, err := operations.TransformToAttribute("request/", req); err == nil {
	// 	span.SetAttributes(kv...)
	// }

	defer span.End()

	// 參數驗證
	var req UpdatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleBindJSON, err))

		return
	}

	// 從path取得id
	id := c.Param("id")

	var input application_role.UpdatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleCopyToInput, err))

		return
	}

	input.Permissions = strings.Join(req.Permissions, ",")
	input.ID = id

	// // 執行UseCase
	data, err := r.u.UpdateRole(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - UpdateRole")
		http.HandleErrorResponse(c, domainerrors.Wrap(ErrorCodeRoleUsecase, err))

		return
	}

	http.SuccessResponse(c, NewResponse(data))
}
