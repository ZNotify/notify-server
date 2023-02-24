package middleware

import (
	"github.com/gin-gonic/gin"
	"notify-api/server/middleware/serverTiming"
)

const timingContextKey = "Timing-Context"

type ServerTimingWriter struct {
	gin.ResponseWriter
	timing          *serverTiming.Timing
	fullCycleMetric *serverTiming.Metric
}

func (w ServerTimingWriter) WriteHeader(code int) {
	w.fullCycleMetric.Stop()
	w.Header().Set(serverTiming.HeaderKey, w.timing.String())
	w.ResponseWriter.WriteHeader(code)
}

func ServerTiming(c *gin.Context) {
	timing := &serverTiming.Timing{}

	c.Writer = &ServerTimingWriter{
		c.Writer,
		timing,
		timing.NewMetric("FullCycle").Start(),
	}

	c.Set(timingContextKey, timing)

	c.Next()
}

func GetTiming(c *gin.Context) *serverTiming.Timing {
	return c.MustGet(timingContextKey).(*serverTiming.Timing)
}
