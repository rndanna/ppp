package graph

import "gateway-service/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	gatewayUS usecase.GatewayUSI
}

func NewResolver(us usecase.GatewayUSI) *Resolver {
	return &Resolver{
		gatewayUS: us,
	}
}
