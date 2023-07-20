package operations

import (
	"errors"
)

var (
	// GoogleCloudOperation -.
	ErrJaegerNotSupported  = errors.New("jaeger is not supported yet")
	ErrBatcherNotSupported = errors.New("batcher is not supported")
)

// InitNewTracer -.
func InitNewTracer(host string, _ int, batcher string, sampleRate float64, enabled bool) error {
	if !enabled {
		return nil
	}

	switch batcher {
	case "gcp":
		GoogleCloudOperationInit(host, sampleRate)
	case "jaeger":
		return ErrJaegerNotSupported
	default:
		return ErrBatcherNotSupported
	}

	return nil
}
