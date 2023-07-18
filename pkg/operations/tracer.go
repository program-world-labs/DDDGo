package operations

import (
	"errors"
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
		return errors.New("jaeger is not supported yet")
	default:
		return errors.New("batcher is not supported")
	}

	return nil
}
