package wallet

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/program-world-labs/DDDGo/internal/adapter/http"
	application_wallet "github.com/program-world-labs/DDDGo/internal/application/wallet"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

type walletRoutes struct {
	u application_wallet.IService
	l pwlogger.Interface
}

func NewWalletRoutes(handler *gin.RouterGroup, u application_wallet.IService, l pwlogger.Interface) {
	r := &walletRoutes{u, l}

	h := handler.Group("/wallet")
	{
		h.POST("/create", r.create)
		h.GET("/list", r.list)
		h.GET("/detail/:id", r.detail)
		h.PUT("/update/:id", r.update)
		h.DELETE("/delete/:id", r.delete)
		// h.PUT("/assign-wallet/:id", r.assignWallet)
	}
}

// @Summary     Create wallet
// @Description Create wallet
// @ID          CreateWallet
// @Tags  	    Wallet
// @Accept      json
// @Produce     json
// @Param		body	body		CreatedRequest	true	"Wallet Create Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/wallet/create [post].
//
//nolint:dupl // business logic is different
func (r *walletRoutes) create(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-create-wallet")

	defer span.End()

	// 參數驗證
	var req CreatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletBindJSON, err, span))

		return
	}

	// // 執行UseCase
	var input application_wallet.CreatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletCopyToInput, err, span))

		return
	}

	walletEntity, err := r.u.CreateWallet(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - CreateWallet")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(walletEntity))
}

// @Summary     List wallet
// @Description List wallet
// @ID          ListWallet
// @Tags  	    Wallet
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
// @Router			/wallet/list [get].
//
//nolint:dupl // business logic is different
func (r *walletRoutes) list(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-list-wallet")

	defer span.End()

	// 參數驗證
	var req ListGotRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindQuery")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletBindQuery, err, span))

		return
	}

	var input application_wallet.ListGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletCopyToInput, err, span))

		return
	}

	// // 執行UseCase
	walletEntities, err := r.u.GetWalletList(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - ListWallet")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponseList(walletEntities))
}

// @Summary     Detail wallet
// @Description Detail wallet
// @ID          DetailWallet
// @Tags  	    Wallet
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/wallet/detail/{id} [get].
//
//nolint:dupl // business logic is different
func (r *walletRoutes) detail(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-detail-wallet")

	defer span.End()

	// 參數驗證
	var req DetailGotRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletBindQuery, err, span))

		return
	}

	var input application_wallet.DetailGotInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletCopyToInput, err, span))

		return
	}

	// 執行UseCase
	walletEntity, err := r.u.GetWalletDetail(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DetailWallet")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(walletEntity))
}

// @Summary     Update wallet
// @Description Update wallet
// @ID          UpdateWallet
// @Tags  	    Wallet
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Param		body	body		UpdatedRequest	true	"Wallet Update Request"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/wallet/update/{id} [put].
func (r *walletRoutes) update(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-update-wallet")

	defer span.End()

	// 參數驗證
	var req UpdatedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindJSON")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletBindJSON, err, span))

		return
	}

	// 從path取得id
	id := c.Param("id")

	var input application_wallet.UpdatedInput
	err := copier.Copy(&input, &req)

	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletCopyToInput, err, span))

		return
	}

	input.ID = id

	// // 執行UseCase
	data, err := r.u.UpdateWallet(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - UpdateWallet")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(data))
}

// @Summary     Delete wallet
// @Description Delete wallet
// @ID          DeleteWallet
// @Tags  	    Wallet
// @Accept      json
// @Produce     json
// @Param		id	path	string	true	"ID"
// @Success		200		{object}	http.Response{data=Response}
// @Failure		400		{object}	http.Response
// @Failure		500		{object}	http.Response
// @Router			/wallet/delete/{id} [delete].
//
//nolint:dupl // business logic is different
func (r *walletRoutes) delete(c *gin.Context) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(c.Request.Context(), "adapter-delete-wallet")

	defer span.End()

	// 參數驗證
	var req DeletedRequest
	if err := c.ShouldBindUri(&req); err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("ShouldBindUri")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletBindQuery, err, span))

		return
	}

	var input application_wallet.DeletedInput

	err := copier.Copy(&input, &req)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Copy")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletCopyToInput, err, span))

		return
	}

	// 執行UseCase
	info, err := r.u.DeleteWallet(ctx, &input)
	if err != nil {
		r.l.Error().Object("Adapter", ErrorEvent{err}).Msg("Usecase - DeleteWallet")
		http.HandleErrorResponse(c, domainerrors.WrapWithSpan(ErrorCodeWalletUsecase, err, span))

		return
	}

	span.SetStatus(codes.Ok, "OK")
	http.SuccessResponse(c, NewResponse(info))
}
