package controller

import (
	"net/http"

	"notify-api/serve/types"
)

// Alive godoc
// @Summary     Server Heartbeat
// @Description Check if the server is alive
// @Produce     plain
// @Success     204 {string} string ""
// @Router      /alive [get]
func Alive(context *types.Ctx) {
	context.String(http.StatusNoContent, "")
}