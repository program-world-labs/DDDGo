package currency

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/program-world-labs/DDDGo/internal/adapter/http"
	application_currency "github.com/program-world-labs/DDDGo/internal/application/currency"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

type currencyRoutes struct {
	u application_currency.IService
	l pwlogger.Interface
}

func NewCurrencyRoutes(handler *gin.RouterGroup, u application_currency.IService, l pwlogger.Interface) {
	r := &currencyRoutes{u, l}

	h := handler.Group("/currency")
	{
		h.POST("/create", r.create)
		h.GET("/list", r.list)
		h.GET("/detail/:id", r.detail)
		h.PUT("/update/:id", r.update)
		h.DELETE("/delete/:id", r.delete)
		// h.PUT("/assign-currency/:id", r.assignCurrency)
	}
}

// @Summary     Create currency
// @Description Create currency
// @ID          CreateCurrency
// @Tags  	    Currency
// @Accept      json
// @Produce     json
// @Param		body	body		CreatedRequest	true	"Currency Create Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/currency/create [post].
//
//nolint:dupl // business logic is different
func (r *currencyRoutes) create(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-create-currency")

	defer span.End()

	// 參數驗證
	var req CreatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeBindJSON, err, span))

		return
	}

	// // 執行UseCase
	var input application_currency.CreatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeCopyToInput, err, span))

		return
	}

	currencyEntity, err := r.u.CreateCurrency(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - CreateCurrency")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeExecuteUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(currencyEntity))
}

// @Summary     List currency
// @Description List currency
// @ID          ListCurrency
// @Tags  	    Currency
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
// @Router			/currency/list [get].
//
//nolint:dupl // business logic is different
func (r *currencyRoutes) list(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-list-currency")

	defer span.End()

	// 參數驗證
	var req ListGotRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindQuery")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeBindQuery, err, span))

		return
	}

	var input application_currency.ListGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeCopyToInput, err, span))

		return
	}

	// // 執行UseCase
	currencyEntities, err := r.u.GetCurrencyList(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - ListCurrency")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeExecuteUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponseList(currencyEntities))
}

// @Summary     Detail currency
// @Description Detail currency
// @ID          DetailCurrency
// @Tags  	    Currency
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/currency/detail/{id} [get].
//
//nolint:dupl // business logic is different
func (r *currencyRoutes) detail(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-detail-currency")

	defer span.End()

	// 參數驗證
	var req DetailGotRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeBindQuery, err, span))

		return
	}

	var input application_currency.DetailGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeCopyToInput, err, span))

		return
	}

	// 執行UseCase
	currencyEntity, err := r.u.GetCurrencyDetail(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DetailCurrency")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeExecuteUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(currencyEntity))
}

// @Summary     Update currency
// @Description Update currency
// @ID          UpdateCurrency
// @Tags  	    Currency
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Param		body	body		UpdatedRequest	true	"Currency Update Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/currency/update/{id} [put].
func (r *currencyRoutes) update(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-update-currency")

	defer span.End()

	// 參數驗證
	var req UpdatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeBindJSON, err, span))

		return
	}

	// 從path取得id
	id := c.Param("id")

	var input application_currency.UpdatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeCopyToInput, err, span))

		return
	}

	input.ID = id

	// // 執行UseCase
	data, err := r.u.UpdateCurrency(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - UpdateCurrency")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeExecuteUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(data))
}

// @Summary     Delete currency
// @Description Delete currency
// @ID          DeleteCurrency
// @Tags  	    Currency
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/currency/delete/{id} [delete].
//
//nolint:dupl // business logic is different
func (r *currencyRoutes) delete(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-delete-currency")

	defer span.End()

	// 參數驗證
	var req DeletedRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeBindQuery, err, span))

		return
	}

	var input application_currency.DeletedInput

	err := copier.Copy(&input, &req)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeCopyToInput, err, span))

		return
	}

	// 執行UseCase
	info, err := r.u.DeleteCurrency(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DeleteCurrency")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeExecuteUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(info))
}
