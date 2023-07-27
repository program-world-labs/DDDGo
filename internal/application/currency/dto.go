package currency

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/program-world-labs/DDDGo/internal/application/utils"
	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
)

type CreatedInput struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (c *CreatedInput) Validate() error {
	validate := validator.New()

	// 提取錯誤訊息
	err := validate.Struct(c)
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Currency", e)
		}

		return err
	}

	return nil
}

func (c *CreatedInput) ToEntity() *entity.Currency {
	return &entity.Currency{
		Name:   c.Name,
		Symbol: c.Symbol,
	}
}

type UpdatedInput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (c *UpdatedInput) Validate() error {
	validate := validator.New()

	// 提取錯誤訊息
	err := validate.Struct(c)
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Currency", e)
		}

		return err
	}

	return nil
}

func (c *UpdatedInput) ToEntity() *entity.Currency {
	return &entity.Currency{
		ID:     c.ID,
		Name:   c.Name,
		Symbol: c.Symbol,
	}
}

type DeletedInput struct {
	ID string `json:"id" validate:"required,alphanum,len=20"`
}

func (i *DeletedInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Currency", e)
		}

		return err
	}

	return nil
}

type DetailGotInput struct {
	ID string `json:"id" validate:"required,alphanum,len=20"`
}

func (i *DetailGotInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Currency", e)
		}

		return err
	}

	return nil
}

type ListGotInput struct {
	Limit      int      `json:"limit"`
	Offset     int      `json:"offset"`
	FilterName string   `json:"filterName" validate:"omitempty,oneof=name"`
	SortFields []string `json:"sortFields" validate:"dive,oneof=id name updated_at created_at"`
	Dir        string   `json:"dir" validate:"oneof=asc desc"`
}

func (i *ListGotInput) Validate() error {
	err := validator.New().Struct(i)
	// 提取錯誤訊息
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return utils.HandleValidationError("Currency", e)
		}

		return err
	}

	return nil
}

func (i *ListGotInput) ToSearchQuery() *domain.SearchQuery {
	sq := &domain.SearchQuery{
		Page: domain.Page{
			Limit:  i.Limit,
			Offset: i.Offset,
		},
	}
	if i.FilterName != "" {
		sq.Filters = append(sq.Filters, domain.Filter{
			FilterField: "name",
			Operator:    "like",
			Value:       i.FilterName,
		})
	}

	if len(i.SortFields) > 0 {
		for _, sortField := range i.SortFields {
			sq.Orders = append(sq.Orders, domain.Order{
				OrderField: sortField,
				Dir:        i.Dir,
			})
		}
	}

	return sq
}

func (i *ListGotInput) ToEntity() *entity.Currency {
	return &entity.Currency{}
}

type Output struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Symbol         string          `json:"symbol"`
	WalletBalances []WalletBalance `json:"walletBalances" `
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
	DeletedAt      time.Time       `json:"deletedAt"`
}

type WalletBalance struct {
	ID         string `json:"id"`
	WalletID   string `json:"walletId"`
	CurrencyID string `json:"currencyId"`
	Balance    uint   `json:"balance"`
	Decimal    uint   `json:"decimal"`
}

type OutputList struct {
	Offset int64    `json:"offset"`
	Limit  int64    `json:"limit"`
	Total  int64    `json:"total"`
	Items  []Output `json:"items"`
}

func NewOutput(e *entity.Currency) *Output {
	var walletBalances []WalletBalance
	for _, v := range e.WalletBalances {
		walletBalances = append(walletBalances, WalletBalance{
			ID:         v.ID,
			WalletID:   v.WalletID,
			CurrencyID: v.CurrencyID,
			Balance:    v.Balance,
			Decimal:    v.Decimal,
		})
	}

	return &Output{
		ID:             e.ID,
		Name:           e.Name,
		Symbol:         e.Symbol,
		WalletBalances: walletBalances,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
		DeletedAt:      e.DeletedAt,
	}
}

func NewListOutput(e *domain.List) *OutputList {
	var outputList []Output
	for _, v := range e.Data {
		outputList = append(outputList, *NewOutput(v.(*entity.Currency)))
	}

	return &OutputList{
		Offset: e.Offset,
		Limit:  e.Limit,
		Total:  e.Total,
		Items:  outputList,
	}
}
