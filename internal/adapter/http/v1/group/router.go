package group

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/program-world-labs/DDDGo/internal/adapter/http"
	application_group "github.com/program-world-labs/DDDGo/internal/application/group"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

type groupRoutes struct {
	u application_group.IService
	l pwlogger.Interface
}

func NewGroupRoutes(handler *gin.RouterGroup, u application_group.IService, l pwlogger.Interface) {
	r := &groupRoutes{u, l}

	h := handler.Group("/group")
	{
		h.POST("/create", r.create)
		h.GET("/list", r.list)
		h.GET("/detail/:id", r.detail)
		h.PUT("/update/:id", r.update)
		h.DELETE("/delete/:id", r.delete)
		// h.PUT("/assign-group/:id", r.assignGroup)
	}
}

// @Summary     Create group
// @Description Create group
// @ID          CreateGroup
// @Tags  	    Group
// @Accept      json
// @Produce     json
// @Param		body	body		CreatedRequest	true	"Group Create Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/group/create [post].
func (r *groupRoutes) create(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-create-group")

	defer span.End()

	// 參數驗證
	var req CreatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupBindJSON, err, span))

		return
	}

	// // 執行UseCase
	var input application_group.CreatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupCopyToInput, err, span))

		return
	}

	groupEntity, err := r.u.CreateGroup(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - CreateGroup")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(groupEntity))
}

// @Summary     List group
// @Description List group
// @ID          ListGroup
// @Tags  	    Group
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
// @Router			/group/list [get].
//
//nolint:dupl // business logic is different
func (r *groupRoutes) list(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-list-group")

	defer span.End()

	// 參數驗證
	var req ListGotRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindQuery")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupBindQuery, err, span))

		return
	}

	var input application_group.ListGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupCopyToInput, err, span))

		return
	}

	// // 執行UseCase
	groupEntities, err := r.u.GetGroupList(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - ListGroup")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponseList(groupEntities))
}

// @Summary     Detail group
// @Description Detail group
// @ID          DetailGroup
// @Tags  	    Group
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/group/detail/{id} [get].
//
//nolint:dupl // business logic is different
func (r *groupRoutes) detail(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-detail-group")

	defer span.End()

	// 參數驗證
	var req DetailGotRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupBindQuery, err, span))

		return
	}

	var input application_group.DetailGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupCopyToInput, err, span))

		return
	}

	// 執行UseCase
	groupEntity, err := r.u.GetGroupDetail(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DetailGroup")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(groupEntity))
}

// @Summary     Update group
// @Description Update group
// @ID          UpdateGroup
// @Tags  	    Group
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Param		body	body		UpdatedRequest	true	"Group Update Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/group/update/{id} [put].
func (r *groupRoutes) update(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-update-group")

	defer span.End()

	// 參數驗證
	var req UpdatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupBindJSON, err, span))

		return
	}

	// 從path取得id
	id := c.Param("id")

	var input application_group.UpdatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupCopyToInput, err, span))

		return
	}

	input.ID = id

	// // 執行UseCase
	data, err := r.u.UpdateGroup(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - UpdateGroup")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(data))
}

// @Summary     Delete group
// @Description Delete group
// @ID          DeleteGroup
// @Tags  	    Group
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/group/delete/{id} [delete].
//
//nolint:dupl // business logic is different
func (r *groupRoutes) delete(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-delete-group")

	defer span.End()

	// 參數驗證
	var req DeletedRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupBindQuery, err, span))

		return
	}

	var input application_group.DeletedInput

	err := copier.Copy(&input, &req)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupCopyToInput, err, span))

		return
	}

	// 執行UseCase
	info, err := r.u.DeleteGroup(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DeleteGroup")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeGroupUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(info))
}
