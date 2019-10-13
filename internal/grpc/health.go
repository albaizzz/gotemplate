package apigrpc

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

// HealthImpl is implementation of grpc healthimpl
type HealthImpl struct{}

// Check implements grpc healthcheck
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// Watch implement interface of healthchek
// unusable right now
func (h *HealthImpl) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return nil
}
