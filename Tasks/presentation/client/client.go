package client

import (
	"github.com/go-kit/kit/endpoint"
)

type RestClient interface {
	MakeEndpoint(apiPath string, method string) endpoint.Endpoint
}
