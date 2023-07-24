package wallet

import (
	"github.com/rs/zerolog"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeWalletUsecase = domainerrors.ErrorCodeAdapterHTTPWallet + iota
	ErrorCodeWalletBindJSON
	ErrorCodeWalletCopyToInput
	ErrorCodeWalletBindQuery
	ErrorCodeWalletValidateInput
	ErrorCodeWalletBindURI
)

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "WalletError")
}
