package handler

import (
	"github.com/DeniesKresna/brigunaduty/service/extensions/terror"
	"github.com/DeniesKresna/brigunaduty/types/models"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) AuthLogin(ctx *gin.Context) {
	var (
		loginReq models.LoginRequest
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	authResp, terr := h.UserUsecase.AuthLogin(ctx, loginReq.Email, loginReq.Password)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, authResp)
	return
}
