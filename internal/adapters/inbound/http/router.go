package httpadapter

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(logger *slog.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		requestIDMiddleware(),
		loggingMiddleware(logger),
		recoveryMiddleware(logger),
		corsMiddleware(),
	)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/ready", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "ready",
		})
	})

	return router
}
