package handler

import (
	"github.com/DeniesKresna/brigunaduty/service/extensions/terror"
	"github.com/DeniesKresna/brigunaduty/types/models"
	"github.com/DeniesKresna/gohelper/utint"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) UserGetByID(ctx *gin.Context) {
	var (
		id   int64
		terr terror.ErrInterface
	)

	id = utint.Convert64FromString(ctx.Param("id"), 0)

	car, terr := h.UserUsecase.UserGetByID(ctx, id)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, car)
	return
}

func (h UserHandler) UserGetByEmail(ctx *gin.Context) {
	var (
		emailReq models.EmailRequest
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&emailReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	user, terr := h.UserUsecase.UserGetByEmail(ctx, emailReq.Email)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, user)
	return
}

func (h UserHandler) UserCreate(ctx *gin.Context) {
	var (
		user models.User
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	user, terr = h.UserUsecase.UserCreate(ctx, user)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, user)
	return
}

func (h UserHandler) UserSearch(ctx *gin.Context) {
	var (
		search models.DbSearchObject
		terr   terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&search); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.UserUsecase.UserSearch(ctx, search)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
	return
}

func (h UserHandler) UserUpdate(ctx *gin.Context) {
	var (
		user models.User
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	user, terr = h.UserUsecase.UserUpdate(ctx, user)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, user)
	return
}

func (h UserHandler) UserGetAllEmployee(ctx *gin.Context) {
	var (
		user models.UserRole
		terr terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	res, terr := h.UserUsecase.UserGetAllEmployee(ctx, user.Name)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, res)
	return
}
