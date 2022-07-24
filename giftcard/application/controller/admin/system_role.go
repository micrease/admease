package admin

import (
	"giftcard/application/controller"
	"giftcard/application/dto"
	"giftcard/application/service"
	"giftcard/library/context/api"
	"giftcard/library/context/result"
	"giftcard/library/validate"
)

type SystemRole struct {
	controller.BaseController
}

func (this SystemRole) Save(ctx *api.Context) *result.Result {
	var req dto.SystemRoleSaveReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemRole()
	err := service.Save(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (this SystemRole) Update(ctx *api.Context) *result.Result {
	var req dto.SystemRoleSaveReq
	validate.BindWithPanic(ctx, &req)
	service := service.NewSystemRole()
	err := service.Save(ctx, req)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success()
}

func (this SystemRole) Index(ctx *api.Context) *result.Result {
	resp, err := service.NewSystemRole().GetPageList(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}

func (this SystemRole) List(ctx *api.Context) *result.Result {
	resp, err := service.NewSystemRole().GetList(ctx)
	if err != nil {
		return result.ServerError(err)
	}
	return result.Success(resp)
}
