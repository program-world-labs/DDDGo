package currency

import (
	"context"
)

type IService interface {
	// Command
	CreateCurrency(ctx context.Context, currencyInfo *CreatedInput) (*Output, error)
	// AssignCurrency(ctx context.Context, currencyInfo *AssignedInput) (*Output, error)
	UpdateCurrency(ctx context.Context, currencyInfo *UpdatedInput) (*Output, error)
	DeleteCurrency(ctx context.Context, currencyInfo *DeletedInput) (*Output, error)
	// // Query
	GetCurrencyList(ctx context.Context, currencyInfo *ListGotInput) (*OutputList, error)
	GetCurrencyDetail(ctx context.Context, currencyInfo *DetailGotInput) (*Output, error)
}
