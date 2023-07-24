package wallet

import (
	"context"
)

type IService interface {
	// Command
	CreateWallet(ctx context.Context, walletInfo *CreatedInput) (*Output, error)
	UpdateWallet(ctx context.Context, walletInfo *UpdatedInput) (*Output, error)
	DeleteWallet(ctx context.Context, walletInfo *DeletedInput) (*Output, error)
	// // Query
	GetWalletList(ctx context.Context, walletInfo *ListGotInput) (*OutputList, error)
	GetWalletDetail(ctx context.Context, walletInfo *DetailGotInput) (*Output, error)
}
