package dto

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeDtoBase = domainerrors.ErrorCodeInfraDTO + domainerrors.ErrorCodeInfraDTO + domainerrors.ErrorCodeInfraDTOBase + iota
	ErrorCodeTransform
	ErrorCodeBackToDomain
	ErrorCodeToJSON
	ErrorCodeDecodeJSON
	ErrorCodeInvalidFilterField
	ErrorCodeInvalidOrderField
	ErrorCodeParseMap
)

var (
	ErrParesMapFailed = errors.New("parse map failed")
	ErrCastTypeFailed = errors.New("cast type failed")
)
