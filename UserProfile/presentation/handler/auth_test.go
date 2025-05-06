package handler

import (
	"context"
	"knovel/userprofile/application/auth"
	"net/http"
	httptest "net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	scenarios := []struct {
		name         string
		expectStatus int
		contextFunc  func() (*gin.Context, *httptest.ResponseRecorder)
		authHandler  func(c *gomock.Controller, ctx context.Context) auth.Authentication
	}{
		{
			name: "success case",
			contextFunc: func() (*gin.Context, *httptest.ResponseRecorder) {
				gin.SetMode(gin.TestMode)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = &http.Request{
					Header: make(http.Header),
				}
				c.Request.Header.Add("x-auth-user-id", "628dec27ac366d22d62446d3")
				c.Request.URL = &url.URL{}
				return c, w
			},
			authHandler: func(c *gomock.Controller, ctx context.Context) auth.Authentication {
				m := auth.NewMockAuthentication(c)
				m.EXPECT().SignIn(ctx, &auth.SignInRequest{
					Email:    "test@test.com",
					PassWord: "123",
				}).Return(auth.ApplicationResponse{}, nil)
				return m
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctx, w := scenario.contextFunc()
			h := &AuthController{
				authen: scenario.authHandler(ctrl, ctx),
			}
			h.SignIn(ctx)
			assert.Equal(t, scenario.expectStatus, w.Code)
		})
	}
}
