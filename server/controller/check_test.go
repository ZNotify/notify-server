package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"notify-api/utils/config"

	"notify-api/utils/user"

	"github.com/gin-gonic/gin"

	"notify-api/server/types"
)

func TestCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Load("test_config")
	user.Init()
	t.Run("check success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/check?user_id=test", nil)
		types.WrapHandler(Check)(c)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		if w.Body.String() != "{\"code\":200,\"body\":true}" {
			t.Errorf("Expected body %s, got %s", "true", w.Body.String())
		}
	})

	t.Run("check fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/check?user_id=error", nil)
		types.WrapHandler(Check)(c)
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		if w.Body.String() != "{\"code\":200,\"body\":false}" {
			t.Errorf("Expected body %s, got %s", "false", w.Body.String())
		}
	})
}