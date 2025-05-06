package application

import (
	"context"
	"knovel/userprofile/application/auth"
	"knovel/userprofile/application/pkg"
	"knovel/userprofile/domain/repositories"
	"net/http"
	"strings"
)

type Application struct {
	authRepository repositories.RoleRepository
	userRepository repositories.UserProfileRepository
}

func NewApplication(authRepository repositories.RoleRepository, userRepository repositories.UserProfileRepository) *Application {
	return &Application{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

var _ auth.Authentication = (*Application)(nil)
var _ auth.Authorization = (*Application)(nil)

func (app *Application) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.ClaimInfo, error) {
	user, err := app.userRepository.SignIn(ctx, req.Email, req.PassWord)
	if err != nil {
		return nil, err
	}

	userRoles, err := app.authRepository.GetRolesOfUser(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	roles := make([]string, 0)
	for _, userRole := range userRoles {
		if userRole.DeletedAt == nil {
			roles = append(roles, userRole.Name)
		}
	}
	claimInfo := &auth.ClaimInfo{
		Id:        user.Id,
		Email:     req.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      roles,
		Code:      pkg.GenerateRandom(20),
	}

	return claimInfo, nil
}

func (app *Application) Authorize(ctx context.Context, req auth.AuthorizeInfo) (auth.ApplicationResponse, error) {
	user, err := app.userRepository.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return auth.ApplicationResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, err
	}
	if user.DeletedAt != nil {
		return auth.ApplicationResponse{
			Status:  http.StatusBadRequest,
			Message: "user is deleted",
		}, nil
	}

	userRoles, err := app.authRepository.GetRolesOfUser(ctx, user.Id)
	if err != nil {
		return auth.ApplicationResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, err
	}

	roles := make([]string, 0)
	roleIds := make([]int, 0)
	for _, userRole := range userRoles {
		if userRole.DeletedAt == nil {
			roles = append(roles, userRole.Name)
			roleIds = append(roleIds, userRole.Id)
		}
	}

	//compare roles with claim's roles
	if !pkg.CompareArrays(req.Roles, roles) {
		return auth.ApplicationResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		}, nil
	}

	if strings.Trim(req.Permission, " ") != "" {
		// get list permissions base on user's roles to find out user has permission to excute a function or not
		userPermissions, err := app.authRepository.GetPermissionByRoles(ctx, roleIds)
		if err != nil {
			return auth.ApplicationResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}, err
		}
		if len(userPermissions) == 0 {
			return auth.ApplicationResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			}, nil
		}
		isAuthorized := false
		for _, entity := range userPermissions {
			if entity.Name == req.Permission && entity.SrvName == req.ServiceName && entity.FuncName == req.FuncName {
				isAuthorized = true
				break
			}
		}
		if !isAuthorized {
			return auth.ApplicationResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
			}, nil
		}
	}

	return auth.ApplicationResponse{
		Status:  http.StatusOK,
		Message: "authorized",
	}, nil
}

// FetchPermissionsOfService implements auth.Authorization.
func (app *Application) FetchPermissionsOfService(ctx context.Context, serviceName string) ([]*auth.PermissionOfService, error) {
	//
	permissions, err := app.authRepository.GetPermissionOfService(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	permissionsOfService := make([]*auth.PermissionOfService, 0)
	for _, permission := range permissions {
		permissionsOfService = append(permissionsOfService, &auth.PermissionOfService{
			FuncName:   permission.FuncName,
			Permission: permission.Name,
		})
	}
	return permissionsOfService, nil
}
