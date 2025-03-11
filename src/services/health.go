package services

type HealthCheckResponse struct {
	Check   bool   `json:"check"`
	Version string `json:"version"`
}

func HealthCheck() HealthCheckResponse {
	response := HealthCheckResponse{
		Check:   true,
		Version: "1.0.0",
	}

	return response
}
