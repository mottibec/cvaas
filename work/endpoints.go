package work

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListEndpoint   endpoint.Endpoint
	CountEndpoint  endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	TagsEndpoint   endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}
