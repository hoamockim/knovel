package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"knovel/tasks/infrastructure/config"
	"knovel/tasks/presentation/client"
	presentation "knovel/tasks/presentation/dto"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

const (
	JsonContentType = "application/json"
)

type RestClient struct {
	ContentType  string
	ClientPath   string
	ApiKey       string
	ResponseData interface{}
}

func NewAuthorClient() client.RestClient {
	authorizeConfig := config.GetAuthorizeClient()
	return &RestClient{
		ContentType:  JsonContentType,
		ClientPath:   authorizeConfig.Path,
		ApiKey:       authorizeConfig.Key,
		ResponseData: &presentation.AuthorizeRespone{},
	}
}

func (client *RestClient) MakeEndpoint(apiPath string, method string) endpoint.Endpoint {
	fullURL, err := url.Parse(fmt.Sprintf("%v/%v", client.ClientPath, apiPath))

	if err != nil {

	}
	return httptransport.NewClient(
		method, fullURL,
		client.encodeRequest,
		client.decodeResponse,
	).Endpoint()
}

func (client *RestClient) encodeRequest(ctx context.Context, r *http.Request, reqBody interface{}) error {
	r.Header.Add("Content-Type", client.ContentType)
	r.Header.Add("X-Api-Key", client.ApiKey)

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqBody)
	if err != nil {
		return err
	}
	r.Body = io.NopCloser(&buf)

	return nil
}

func (client *RestClient) decodeResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	if err := json.NewDecoder(r.Body).Decode(client.ResponseData); err != nil {
		return nil, err
	}
	return client.ResponseData, nil
}
