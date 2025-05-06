package handler

import (
	"context"
	"knovel/userprofile/application"
	"knovel/userprofile/application/auth"
	"knovel/userprofile/presentation/dto"
	"knovel/userprofile/presentation/util/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignIn(ctx *gin.Context)
	Authorize(ctx *gin.Context)
	FetchPermissionOfService(ctx *gin.Context)
}

type AuthController struct {
	authen auth.Authentication
	author auth.Authorization
}

var _ AuthHandler = (*AuthController)(nil)

func NewAuthHandler(app *application.Application) AuthHandler {
	return &AuthController{
		authen: app,
		author: app,
	}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var req dto.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	claim, err := c.authen.SignIn(context.Background(), &auth.SignInRequest{
		Email:    req.Email,
		PassWord: req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	token := jwt.GenerateToken(jwt.ClaimInfo{
		Email:     claim.Email,
		Id:        claim.Id,
		FirstName: claim.FirstName,
		LastName:  claim.LastName,
		Role:      claim.Role,
	})
	res := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	ctx.JSON(http.StatusOK, &res)
}

func (c *AuthController) Authorize(ctx *gin.Context) {
	var req dto.AuthorizeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result, err := c.author.Authorize(ctx.Request.Context(), auth.AuthorizeInfo{
		UserId:      req.UserId,
		Roles:       req.Roles,
		Permission:  req.Permission,
		ServiceName: req.ServiceName,
		FuncName:    req.FuncName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	var authorStatus dto.AuthorizeStatus
	switch result.Status {
	case http.StatusUnauthorized:
		authorStatus = dto.UnAuthorized
	case http.StatusOK:
		authorStatus = dto.Authorized
	default:
		authorStatus = dto.Error
	}
	ctx.JSON(result.Status, &dto.AuthorizeRespone{
		Status:  authorStatus,
		Message: result.Message,
	})
}

func (c *AuthController) FetchPermissionOfService(ctx *gin.Context) {
	query, _ := ctx.GetQuery("service_name")
	if query == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "service name is invalid"})
	}
	result, err := c.author.FetchPermissionsOfService(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, &dto.PermissionOfServiceRespone{
		Data: result,
	})
}
