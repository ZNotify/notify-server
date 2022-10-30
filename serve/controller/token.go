package controller

import (
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"notify-api/db/model"
	"notify-api/push"
	"notify-api/serve/types"
	"notify-api/utils"
)

// Token godoc
// @Summary     Create or update token
// @Description Create or update token of device
// @Param       user_id   path     string true  "user_id"
// @Param       device_id path     string true  "device_id should be a valid UUID"
// @Param       channel   formData string true  "channel"
// @Param       token     formData string false "token"
// @Produce     json
// @Success     200 {object} types.Response[bool]
// @Failure     400 {object} types.BadRequestResponse
// @Failure     401 {object} types.UnauthorizedResponse
// @Router      /{user_id}/token/{device_id} [put]
func Token(context *types.Ctx) {
	deviceID := context.Param("device_id")
	if !utils.IsUUID(deviceID) {
		zap.S().Infof("device id %s is not a valid UUID", deviceID)
		context.JSONError(http.StatusBadRequest, errors.New("device_id should be a valid UUID"))
		return
	}

	channel := context.PostForm("channel")
	if !push.Senders.Has(channel) {
		zap.S().Infof("channel %s is not supported", channel)
		context.JSONError(http.StatusBadRequest, errors.New("channel is not supported"))
		return
	}

	token := context.PostForm("token")

	// FIXME: use register mechanism to avoid
	if channel == "WebSocketHost" && token != "" {
		zap.S().Infof("token should be empty for WebSocketHost")
		context.JSONError(http.StatusBadRequest, errors.New("token should be empty for WebSocketHost"))
		return
	}

	err := model.TokenUtils.CreateOrUpdate(context.UserID, deviceID, channel, token)
	if err != nil {
		zap.S().Errorw("create or update token error", "error", err)
		context.JSONError(http.StatusInternalServerError, errors.WithStack(err))
		return
	}
	context.JSONResult(true)
}
