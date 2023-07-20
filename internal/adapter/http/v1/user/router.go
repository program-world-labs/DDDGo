package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/program-world-labs/DDDGo/internal/adapter/http"
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

type userRoutes struct {
	u application_user.IService
	l pwlogger.Interface
}

func NewUserRoutes(handler *gin.RouterGroup, u application_user.IService, l pwlogger.Interface) {
	r := &userRoutes{u, l}

	h := handler.Group("/user")
	{
		h.POST("/create", r.create)
		h.GET("/list", r.list)
		h.GET("/detail/:id", r.detail)
		h.PUT("/update/:id", r.update)
		h.DELETE("/delete/:id", r.delete)
	}
}

// @Summary     Create user
// @Description Create user
// @ID          CreateUser
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Param		body	body		CreatedRequest	true	"User Create Request"
// @Success		200		{object}	http.Response
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/user/create [post].
//
//nolint:dupl // business logic is different
func (r *userRoutes) create(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-create-user")

	defer span.End()

	// 參數驗證
	var req CreatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserBindJSON, err, span))

		return
	}

	// // 執行UseCase
	var input application_user.CreatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserCopyToInput, err, span))

		return
	}

	userEntity, err := r.u.CreateUser(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - CreateUser")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(userEntity))
}

// @Summary     List user
// @Description List user
// @ID          ListUser
// @Tags  	    User
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
// @Router			/user/list [get].
//
//nolint:dupl // business logic is different
func (r *userRoutes) list(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-list-user")

	defer span.End()

	// 參數驗證
	var req ListGotRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindQuery")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserBindQuery, err, span))

		return
	}

	// 執行UseCase
	var input application_user.ListGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserCopyToInput, err, span))

		return
	}

	userList, err := r.u.GetUserList(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - ListUser")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponseList(userList))
}

// @Summary     Detail user
// @Description Detail user
// @ID          DetailUser
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Param		id	path	string		true	"ID"
// @Success		200		{object}	http.Response
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/user/detail/{id} [get].
//
//nolint:dupl // business logic is different
func (r *userRoutes) detail(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-detail-user")

	defer span.End()

	// 參數驗證
	var req DetailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserBindURI, err, span))

		return
	}

	// 執行UseCase
	userEntity, err := r.u.GetUserDetail(ctx, &application_user.DetailGotInput{ID: req.ID})
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DetailUser")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(userEntity))
}

// @Summary     Update user
// @Description Update user
// @ID          UpdateUser
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Param		id	path	string		true	"ID"
// @Param		body	body		UpdateRequest	true	"User Update Request"
// @Success		200		{object}	http.Response
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/user/update/{id} [put].
func (r *userRoutes) update(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-update-user")

	defer span.End()

	// 參數驗證
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserBindJSON, err, span))

		return
	}

	// 取得ID
	id := c.Param("id")

	// 執行UseCase
	var input application_user.UpdatedInput

	err := copier.Copy(&input, &req)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserCopyToInput, err, span))

		return
	}

	input.ID = id

	userEntity, err := r.u.UpdateUser(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - UpdateUser")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(userEntity))
}

// @Summary     Delete user
// @Description Delete user
// @ID          DeleteUser
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Param		id	path	string		true	"ID"
// @Success		200		{object}	http.Response
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/user/delete/{id} [delete].
//
//nolint:dupl // business logic is different
func (r *userRoutes) delete(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-delete-user")

	defer span.End()

	// 參數驗證
	var req DeleteRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserBindURI, err, span))

		return
	}

	// 執行UseCase
	info, err := r.u.DeleteUser(ctx, &application_user.DeletedInput{ID: req.ID})
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DeleteUser")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeUserUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(info))
}
