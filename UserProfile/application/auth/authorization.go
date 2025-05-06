package auth

import "context"

type Authorization interface {
	Authorize(ctx context.Context, authorInfo AuthorizeInfo) (ApplicationResponse, error)
	FetchPermissionsOfService(ctx context.Context, serviceName string) ([]*PermissionOfService, error)
}
