package wallet

import (
	"time"

	application_wallet "github.com/program-world-labs/DDDGo/internal/application/wallet"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
)

type Response struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	Chain          entity.Chain       `json:"chain" binding:"required" example:"Polygon"`
	UserID         string             `json:"userId" binding:"required" example:"abcd-efgh-ijkl-mnop"`
	WalletBalances []BalancesResponse `json:"walletBalances"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	DeletedAt      time.Time          `json:"deletedAt"`
}

type BalancesResponse struct {
	ID         string `json:"id" gorm:"primary_key"`
	WalletID   string `json:"walletId" gorm:"index"`
	CurrencyID string `json:"currencyId" gorm:"index"`
	Balance    uint   `json:"balance"`
	Decimal    uint   `json:"decimal"`
}

type ResponseList struct {
	Offset int64      `json:"offset"`
	Limit  int64      `json:"limit"`
	Total  int64      `json:"total"`
	Items  []Response `json:"items"`
}

func NewResponse(model *application_wallet.Output) Response {
	return Response{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Chain:       model.Chain,
		UserID:      model.UserID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		DeletedAt:   model.DeletedAt,
	}
}

func NewResponseList(modelList *application_wallet.OutputList) ResponseList {
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
