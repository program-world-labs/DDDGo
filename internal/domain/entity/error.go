package entity

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeDomainRole  = domainerrors.ErrorCodeDomainRole + iota // 100000
	ErrorCodeCastToEvent                                           // 100001
	ErrorCodeCast                                                  // 100002
	ErrorCodeDomainUser = domainerrors.ErrorCodeDomainRole
)

var (
	ErrCastToEventFailed = errors.New("cast to event failed")
)
