package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"notify-api/db/model"
	. "notify-api/push"
	pushTypes "notify-api/push/types"
	"notify-api/serve/types"
)

// Send godoc
// @Summary     Send notification
// @Description Send notification to user_id
// @Param       user_id path     string true  "user_id"
// @Param       title   formData string false "title"
// @Param       content formData string true  "content"
// @Param       long    formData string false "long"
// @Produce     json
// @Success     200 {object} types.Response[entity.Message]
// @Failure     400 {object} types.BadRequestResponse
// @Failure     401 {object} types.UnauthorizedResponse
// @Router      /{user_id}/send [post]
// @Router      /{user_id}/send [put]
func Send(context *types.Ctx) {
	// get notification info
	title := context.DefaultPostForm("title", "Notification")
	content := context.PostForm("content")
	long := context.PostForm("long")

	if content == "" {
		context.JSONError(http.StatusBadRequest, "Content can not be empty.")
		return
	}

	pushMsg := &pushTypes.Message{
		ID:        uuid.New().String(),
		UserID:    context.UserID,
		Title:     title,
		Content:   content,
		Long:      long,
		CreatedAt: time.Now(),
	}

	err := Senders.Send(pushMsg)
	if err != nil {
		context.JSONError(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	// Insert message record
	time.Sleep(1 * time.Nanosecond)
	// a trick to generate different timestamp for different message
	// FIXME: use an increasing counter to generate different id
	msg, err := model.MessageUtils.Add(
		pushMsg.ID,
		pushMsg.UserID,
		pushMsg.Title,
		pushMsg.Content,
		pushMsg.Long,
		time.Now())

	if err != nil {
		context.JSONError(http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	context.JSONResult(msg)
	return
}
