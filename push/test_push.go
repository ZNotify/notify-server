//go:build test

package push

import (
	"go.uber.org/zap"

	pushTypes "notify-api/push/entity"
)

func Send(msg *pushTypes.PushMessage) error {
	fields := []zap.Field{
		zap.String("user_id", msg.UserID),
		zap.String("title", msg.Title),
		zap.String("content", msg.Content),
		zap.String("long_content", msg.Long),
		zap.String("priority", string(msg.Priority)),
	}

	zap.L().Info("Try to send message", fields...)
	return nil
}

func Init() {
	activeSenders = availableSenders
}
