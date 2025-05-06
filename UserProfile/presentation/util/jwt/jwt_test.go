package jwt

import (
	"fmt"
	"testing"
)

func Test_GenerateToken(t *testing.T) {
	publicPath := "test_public.key"
	privatePath := "test_private.key"
	InitJWT(publicPath, privatePath)
	claim := ClaimInfo{
		Email: "test@gmail.com",
		Role:  []string{"admin", "user"},
	}
	token := GenerateToken(claim)
	if token == "" {
		t.Error("error")
	}
	fmt.Println(token)
}
