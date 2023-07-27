package currency

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/wallet"
	application_currency "github.com/program-world-labs/DDDGo/internal/application/currency"
)

type Response struct {
	ID             string                    `json:"id"`
	Name           string                    `json:"name"`
	Symbol         string                    `json:"symbol"`
	WalletBalances []wallet.BalancesResponse `json:"walletBalances"`
	CreatedAt      time.Time                 `json:"createdAt"`
	UpdatedAt      time.Time                 `json:"updatedAt"`
	DeletedAt      time.Time                 `json:"deletedAt"`
}

type ResponseList struct {
	Offset int64      `json:"offset"`
	Limit  int64      `json:"limit"`
	Total  int64      `json:"total"`
	Items  []Response `json:"items"`
}

func NewResponse(model *application_currency.Output) Response {
	return Response{
		ID:     model.ID,
		Name:   model.Name,
		Symbol: model.Symbol,
		WalletBalances: func() []wallet.BalancesResponse {
			walletBalances := make([]wallet.BalancesResponse, len(model.WalletBalances))

			for i, v := range model.WalletBalances {
				value := v
				walletBalances[i] = wallet.BalancesResponse{
					ID:         value.ID,
					WalletID:   value.WalletID,
					CurrencyID: value.CurrencyID,
					Balance:    value.Balance,
					Decimal:    value.Decimal,
				}
			}

			return walletBalances
		}(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}

func NewResponseList(modelList *application_currency.OutputList) ResponseList {
	responseList := make([]Response, len(modelList.Items))

	for i, v := range modelList.Items {
		value := v
		responseList[i] = NewResponse(&value)
	}

	return ResponseList{
		Offset: modelList.Offset,
		Limit:  modelList.Limit,
		Total:  modelList.Total,
		Items:  responseList,
	}
}
