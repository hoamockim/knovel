package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SysHandler interface {
	HealthCheck(ctx *gin.Context)
}

type SysHandlerImpl struct {
}

var _ SysHandler = (*SysHandlerImpl)(nil)

func NewSysHandler() SysHandler {
	return &SysHandlerImpl{}
}

func (sys *SysHandlerImpl) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
