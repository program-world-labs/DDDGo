package operations

// InitNewTracer -.
func InitNewTracer(host string, _ int, batcher string, sampleRate float64, enabled bool) error {
	if !enabled {
		return nil
	}

	switch batcher {
	case "gcp":
		GoogleCloudOperationInit(host, sampleRate)
	}

	return nil
}
