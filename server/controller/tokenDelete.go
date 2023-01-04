package controller

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/server/types"
)

// TokenDelete godoc
//
//	@Summary		Delete token
//	@Description	Delete token of device
//	@Param			user_id		path	string	true	"user_id"
//	@Param			device_id	path	string	true	"device_id"
//	@Produce		json
//	@Success		200	{object}	types.Response[bool]
//	@Router			/{user_id}/token/{device_id} [delete]
func TokenDelete(context *types.Ctx) {
	deviceId := context.Param("device_id")

	err := dao.DeviceDao.SafeDeleteDevice(context.UserID, deviceId)
	if err != nil {
		zap.S().Errorw("delete token error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}

	context.JSONResult(true)
}