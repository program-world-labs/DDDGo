package http

import "github.com/program-world-labs/DDDGo/internal/domain/domainerrors"

type AdapterError struct {
	domainerrors.ErrorInfo
}

func NewAdapterError(err error) *AdapterError {
	return &AdapterError{*domainerrors.New(domainerrors.GruopID+domainerrors.ErrorCodeSystem, err.Error())} // A00000000
}
