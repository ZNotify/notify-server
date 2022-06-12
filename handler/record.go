package handler

import (
	"github.com/ZNotify/server/db"
	"github.com/ZNotify/server/db/entity"
	"github.com/ZNotify/server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Record(context *gin.Context) {
	userID, err := utils.RequireAuth(context)
	if err != nil {
		utils.BreakOnError(context, err)
		return
	}

	var messages []entity.Message
	result := db.DB.Where("user_id = ?", userID).
		Where("created_at > ?", time.Now().AddDate(0, 0, -30)).
		Order("created_at desc").
		Find(&messages)
	utils.BreakOnError(context, result.Error)

	var ret []gin.H
	for i := range messages {
		ret = append(ret, messages[i].ToGinH())
	}
	context.JSON(http.StatusOK, ret)
}